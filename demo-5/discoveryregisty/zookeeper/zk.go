package zookeeper

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-zookeeper/zk"
)

const (
	defaultFreq = time.Minute * 30
	timeout     = 5
)

type ZkClient struct {
	// 服务注册专用
	conn      *zk.Conn
	zkServers []string
	scheme    string
	userName  string

	password  string
	timeout   int
	eventChan <-chan zk.Event
	node      string
	ticker    *time.Ticker

	lock    sync.Locker
	options []grpc.DialOption

	resolvers    map[string]*Resolver
	localConns   map[string][]resolver.Address
	balancerName string
}

func (s *ZkClient) Scheme() string { return strings.ToLower(s.scheme) }

type ZkOption func(*ZkClient)

func WithRoundRobin() ZkOption {
	return func(client *ZkClient) {
		client.balancerName = roundrobin.Name
	}
}

func WithUserNameAndPassword(userName, password string) ZkOption {
	return func(client *ZkClient) {
		client.userName = userName
		client.password = password
	}
}

func WithOptions(opts ...grpc.DialOption) ZkOption {
	return func(client *ZkClient) {
		client.options = opts
	}
}

func WithFreq(freq time.Duration) ZkOption {
	return func(client *ZkClient) {
		client.ticker = time.NewTicker(freq)
	}
}

func WithTimeout(timeout int) ZkOption {
	return func(client *ZkClient) {
		client.timeout = timeout
	}
}

func NewClient(zkServers []string, schema string, options ...ZkOption) (*ZkClient, error) {
	client := &ZkClient{
		zkServers:  zkServers,
		scheme:     schema,
		timeout:    timeout,
		localConns: make(map[string][]resolver.Address),
		resolvers:  make(map[string]*Resolver),
		lock:       &sync.Mutex{},
	}
	client.ticker = time.NewTicker(defaultFreq)
	for _, option := range options {
		option(client)
	}
	conn, eventChan, err := zk.Connect(zkServers, time.Duration(client.timeout)*time.Second, zk.WithLogInfo(true))
	if err != nil {
		return nil, err
	}
	if client.userName != "" && client.password != "" {
		if err := conn.AddAuth("digest", []byte(client.userName+":"+client.password)); err != nil {
			return nil, err
		}
	}
	client.eventChan = eventChan
	client.conn = conn
	if err := client.ensureRoot(); err != nil {
		client.CloseZK()
		return nil, err
	}
	resolver.Register(client)
	go client.refresh()
	go client.watch()
	return client, nil
}

func (s *ZkClient) CloseZK() {
	s.conn.Close()
}

func (s *ZkClient) ensureAndCreate(node string) error {
	exists, _, err := s.conn.Exists(node)
	if err != nil {
		return err
	}
	if !exists {
		_, err := s.conn.Create(node, []byte(""), 0, zk.WorldACL(zk.PermAll))
		if err != nil && err != zk.ErrNodeExists {
			return err
		}
	}
	return nil
}

func (s *ZkClient) refresh() {
	for range s.ticker.C {
		fmt.Println("refresh local conns")
		s.lock.Lock()
		for rpcName := range s.resolvers {
			s.flushResolver(rpcName)
		}
		for rpcName := range s.localConns {
			delete(s.localConns, rpcName)
		}
		s.lock.Unlock()
		fmt.Println("refresh local conns success")
	}
}

func (s *ZkClient) flushResolverAndDeleteLocal(serviceName string) {
	fmt.Println("start flush %s", serviceName)
	s.flushResolver(serviceName)
	delete(s.localConns, serviceName)
}

func (s *ZkClient) flushResolver(serviceName string) {
	r, ok := s.resolvers[serviceName]
	if ok {
		r.ResolveNowZK(resolver.ResolveNowOptions{})
	}
}

func (s *ZkClient) GetZkConn() *zk.Conn {
	return s.conn
}

func (s *ZkClient) GetRootPath() string {
	return s.scheme
}

func (s *ZkClient) GetNode() string {
	return s.node
}

func (s *ZkClient) ensureRoot() error {
	return s.ensureAndCreate(s.scheme)
}

func (s *ZkClient) ensureName(rpcRegisterName string) error {
	return s.ensureAndCreate(s.getPath(rpcRegisterName))
}

func (s *ZkClient) getPath(rpcRegisterName string) string {
	return s.scheme + "/" + rpcRegisterName
}

func (s *ZkClient) getAddr(host string, port int) string {
	return net.JoinHostPort(host, strconv.Itoa(port))
}

func (s *ZkClient) AddOption(opts ...grpc.DialOption) {
	s.options = append(s.options, opts...)
}

func (s *ZkClient) GetClientLocalConns() map[string][]resolver.Address {
	return s.localConns
}
