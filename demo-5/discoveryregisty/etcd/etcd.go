package etcd

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"strings"
	"sync"
	"time"
)

const default_conn_TTL = 5

type EtcdRegister struct {
	// 服务注册专用
	cli         *clientv3.Client
	userName    string
	password    string
	etcdAddr    []string
	schema      string
	key         string
	closeCh     chan struct{}
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse

	// 连接超时配置
	timeout int

	// 服务发现专用
	resolvers    map[string]*Resolver
	localConns   map[string][]grpc.ClientConnInterface
	options      []grpc.DialOption
	balancerName string
	// 全局锁操作resolvers 和localConns
	lock sync.Locker
}

// 注意协议名必须小写  grpc v1.55中 url.parse中 会将协议名转为小写
func (r *EtcdRegister) Scheme() string {
	return strings.ToLower(r.schema)
}

func NewClient(etcdAddr []string, schema string, opts ...EtcdOption) (*EtcdRegister, error) {
	register := &EtcdRegister{
		schema:     schema,
		etcdAddr:   etcdAddr,
		timeout:    default_conn_TTL,
		lock:       &sync.Mutex{},
		localConns: make(map[string][]grpc.ClientConnInterface),
		resolvers:  make(map[string]*Resolver),
	}
	for _, opt := range opts {
		opt(register)
	}
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   register.etcdAddr,
		DialTimeout: time.Duration(register.timeout) * time.Second,
		Username:    register.userName,
		Password:    register.password,
	})
	if err != nil {
		panic(err)
	}
	register.cli = client
	resolver.Register(register)
	return register, err
}

type EtcdOption func(*EtcdRegister)

func WithOptions(opts ...grpc.DialOption) EtcdOption {
	return func(client *EtcdRegister) {
		client.options = opts
	}
}

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

func WithTimeout(timeout int) EtcdOption {
	return func(client *EtcdRegister) {
		client.timeout = timeout
	}
}

func GetPrefix(schema, serviceName string) string {
	return fmt.Sprintf("%s:///%s/", schema, serviceName)
}

func (s *EtcdRegister) flushResolverAndDeleteLocal(serviceName string) {
	fmt.Println("start flush ", serviceName)
	s.flushResolver(serviceName)
	delete(s.localConns, serviceName)
}

func (s *EtcdRegister) flushResolver(serviceName string) {
	r, ok := s.resolvers[serviceName]
	if ok {
		r.ResolveNowEtcd(resolver.ResolveNowOptions{})
	}
}

func (r *EtcdRegister) AddOption(opts ...grpc.DialOption) {
	r.options = append(r.options, opts...)
}

func (r *EtcdRegister) GetClientLocalConns() map[string][]grpc.ClientConnInterface {
	return r.localConns
}
