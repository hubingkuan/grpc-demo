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

const Default_TTL = 10

type EtcdRegister struct {
	// 这部分属于服务注册专用
	cli         *clientv3.Client
	userName    string
	password    string
	etcdAddr    []string
	schema      string
	key         string
	closeCh     chan struct{}
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse

	// 全局锁操作resolvers 和localConns
	lock sync.Locker
	// 这部分属于服务发现专用
	resolvers    map[string]*Resolver
	localConns   map[string][]grpc.ClientConnInterface
	options      []grpc.DialOption
	balancerName string
}

// 注意协议名必须小写  url.parse中 会将协议名转为小写
func (r *EtcdRegister) Scheme() string {
	return strings.ToLower(r.schema)
}

func NewClient(etcdAddr []string, schema string, opts ...EtcdOption) (*EtcdRegister, error) {
	register := &EtcdRegister{
		schema:     schema,
		etcdAddr:   etcdAddr,
		lock:       &sync.Mutex{},
		localConns: make(map[string][]grpc.ClientConnInterface),
		resolvers:  make(map[string]*Resolver),
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
	resolver.Register(register)
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

func (r *EtcdRegister) GetClientLocalConns() map[string][]grpc.ClientConnInterface {
	return r.localConns
}
