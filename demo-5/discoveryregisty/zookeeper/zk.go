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
	timeout = 5
)

type ZkClient struct {
	// 服务注册专用
	conn      *zk.Conn
	zkServers []string
	zkRoot    string
	schema    string
	userName  string
	password  string
	// 连接超时时间
	timeout uint64

	node            string
	rpcRegisterName string
	rpcRegisterAddr string

	isStateDisconnected bool
	isRegistered        bool

	// 服务发现专用
	resolvers    map[string]*Resolver
	localConns   map[string][]grpc.ClientConnInterface
	options      []grpc.DialOption
	balancerName string
	// 全局锁
	lock sync.Locker
}

func (s *ZkClient) GetClientLocalConns() map[string][]grpc.ClientConnInterface {
	return s.localConns
}

func (s *ZkClient) Scheme() string { return strings.ToLower(s.schema) }

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

func WithTimeout(timeout uint64) ZkOption {
	return func(client *ZkClient) {
		client.timeout = timeout
	}
}

func NewClient(zkServers []string, schema string, options ...ZkOption) (*ZkClient, error) {
	client := &ZkClient{
		zkServers:  zkServers,
		schema:     schema,
		zkRoot:     "/" + schema,
		timeout:    timeout,
		localConns: make(map[string][]grpc.ClientConnInterface),
		resolvers:  make(map[string]*Resolver),
		lock:       &sync.Mutex{},
	}
	for _, option := range options {
		option(client)
	}
	conn, _, err := zk.Connect(client.zkServers, time.Duration(client.timeout)*time.Second, zk.WithLogInfo(true), zk.WithEventCallback(callback))
	if err != nil {
		return nil, err
	}
	if client.userName != "" && client.password != "" {
		if err := conn.AddAuth("digest", []byte(client.userName+":"+client.password)); err != nil {
			return nil, err
		}
	}
	client.conn = conn
	if err := client.ensureRoot(); err != nil {
		client.CloseZK()
		return nil, err
	}
	resolver.Register(client)
	return client, nil
}

// todo 改变watch逻辑 目前只有服务发现会使用到watch
func callback(e zk.Event) {
	fmt.Println("++++++++++++++++++++++++")
	fmt.Println("path:", e.Path)
	fmt.Println("type:", e.Type.String())
	fmt.Println("state:", e.State.String())
	fmt.Println("------------------------")
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

func (s *ZkClient) flushResolverAndDeleteLocal(serviceName string) {
	fmt.Println("start flush ", serviceName)
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
	return s.zkRoot
}

func (s *ZkClient) GetNode() string {
	return s.node
}

func (s *ZkClient) ensureRoot() error {
	return s.ensureAndCreate(s.zkRoot)
}

func (s *ZkClient) ensureName(rpcRegisterName string) error {
	return s.ensureAndCreate(s.getPath(rpcRegisterName))
}

func (s *ZkClient) getPath(rpcRegisterName string) string {
	return s.zkRoot + "/" + rpcRegisterName
}

func (s *ZkClient) getAddr(host string, port int) string {
	return net.JoinHostPort(host, strconv.Itoa(port))
}

func (s *ZkClient) AddOption(opts ...grpc.DialOption) {
	s.options = append(s.options, opts...)
}
