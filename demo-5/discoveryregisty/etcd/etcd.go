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
)

const Default_TTL = 10

type EtcdClient struct {
	cli         *clientv3.Client
	schema      string
	key         string
	closeCh     chan struct{}
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse
	etcdAddr    []string

	lock    sync.Locker
	options []grpc.DialOption

	resolvers    map[string]*discoveryregisty.Resolver
	localConns   map[string][]resolver.Address
	balancerName string
}

func (s *EtcdClient) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	fmt.Printf("build resolver: %+v, cc: %+v\n", target, cc.UpdateState)
	r := &discoveryregisty.Resolver{}
	r.SetTarget(target)
	r.SetCc(cc)
	r.SetGetConnsRemote(s.GetConnsRemote)
	r.ResolveNowZK(resolver.ResolveNowOptions{})
	s.lock.Lock()
	defer s.lock.Unlock()
	serviceName := strings.TrimLeft(target.URL.Path, "/")
	s.resolvers[serviceName] = r
	fmt.Printf("build resolver finished: %+v, cc: %+v, key: %s\n", target, cc.UpdateState, serviceName)
	return r, nil
}

func (r *EtcdClient) Scheme() string {
	return r.schema
}

type EtcdOption func(*EtcdClient)

func WithRoundRobin() EtcdOption {
	return func(client *EtcdClient) {
		client.balancerName = roundrobin.Name
	}
}

func GetPrefix(schema, serviceName string) string {
	return fmt.Sprintf("%s:///%s/", schema, serviceName)
}

func (r *EtcdClient) AddOption(opts ...grpc.DialOption) {
	r.options = append(r.options, opts...)
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

func (r *EtcdClient) GetClientLocalConns() map[string][]resolver.Address {
	return r.localConns
}
