package etcd

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"grpc-demo/demo-5/discoveryregisty"
	"strings"
	"sync"
	"time"
)

const Default_TTL = 10

type EtcdRegister struct {
	cli         *clientv3.Client
	schema      string
	key         string
	closeCh     chan struct{}
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse
	etcdAddr    []string
	userName    string
	password    string

	lock    sync.Locker
	options []grpc.DialOption

	resolvers    map[string]*discoveryregisty.Resolver
	localConns   map[string][]resolver.Address
	balancerName string
}

func (s *EtcdRegister) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	fmt.Printf("build resolver: %+v\n", target)
	r := &discoveryregisty.Resolver{}
	r.SetTarget(target)
	r.SetCc(cc)
	r.SetGetConnsRemote(s.GetConnsRemote)
	r.ResolveNowZK(resolver.ResolveNowOptions{})
	s.lock.Lock()
	defer s.lock.Unlock()
	serviceName := strings.TrimLeft(target.URL.Path, "/")
	s.resolvers[serviceName] = r
	fmt.Printf("build resolver finished: %+v,key: %s\n", target, serviceName)
	return r, nil
}

func (r *EtcdRegister) Scheme() string {
	return r.schema
}

func NewClient(etcdAddr []string, schema string, opts ...EtcdOption) (*EtcdRegister, error) {
	register := &EtcdRegister{
		schema:     schema,
		etcdAddr:   etcdAddr,
		lock:       &sync.Mutex{},
		resolvers:  make(map[string]*discoveryregisty.Resolver),
		localConns: make(map[string][]resolver.Address),
	}
	for _, opt := range opts {
		opt(register)
	}
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   register.etcdAddr,
		DialTimeout: time.Duration(Default_TTL) * time.Second,
		Username:    register.userName,
		Password:    register.password,
	})
	if err != nil {
		panic(err)
	}
	register.cli = client
	return register, err
}

type EtcdOption func(*EtcdRegister)

func WithRoundRobin() EtcdOption {
	return func(client *EtcdRegister) {
		client.balancerName = roundrobin.Name
	}
}

func WithUserNameAndPassword(userName, password string) EtcdOption {
	return func(client *EtcdRegister) {
		client.userName = userName
		client.password = password
	}
}

func WithOptions(opts ...grpc.DialOption) EtcdOption {
	return func(client *EtcdRegister) {
		client.options = opts
	}
}

func GetPrefix(schema, serviceName string) string {
	return fmt.Sprintf("%s:///%s/", schema, serviceName)
}

func exists(addrList []resolver.Address, addr string) bool {
	for _, v := range addrList {
		if v.Addr == addr {
			return true
		}
	}
	return false
}

func remove(s []resolver.Address, addr string) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}

func (r *EtcdRegister) AddOption(opts ...grpc.DialOption) {
	r.options = append(r.options, opts...)
}

func (r *EtcdRegister) GetClientLocalConns() map[string][]resolver.Address {
	return r.localConns
}
