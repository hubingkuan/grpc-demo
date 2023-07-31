package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

func (r *EtcdRegister) watch(serviceName string) {
	rch := r.cli.Watch(context.Background(), GetPrefix(r.schema, serviceName), clientv3.WithPrefix())
	for n := range rch {
		for _, ev := range n.Events {
			fmt.Printf("etcd event value: %s\n", ev.Kv.Value)
			switch ev.Type {
			case mvccpb.PUT:
				r.lock.Lock()
				r.flushResolverAndDeleteLocal(serviceName)
				r.lock.Unlock()
			case mvccpb.DELETE:
				r.lock.Lock()
				r.flushResolverAndDeleteLocal(serviceName)
				r.lock.Unlock()
			}
		}
	}
}

func (r *EtcdRegister) GetConnsRemote(serviceName string) (conns []resolver.Address, err error) {
	getResponse, err := r.cli.Get(context.Background(), GetPrefix(r.schema, serviceName), clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	for _, kv := range getResponse.Kvs {
		addr := string(kv.Value)
		fmt.Println("get conns from remote", "conn:", addr)
		conns = append(conns, resolver.Address{Addr: addr, ServerName: serviceName})
	}
	return conns, nil
}

func (s *EtcdRegister) GetConns(ctx context.Context, serviceName string, opts ...grpc.DialOption) ([]grpc.ClientConnInterface, error) {
	fmt.Printf("get conns from client, serviceName: %s\n", serviceName)
	s.lock.Lock()
	defer s.lock.Unlock()
	opts = append(s.options, opts...)
	conns := s.localConns[serviceName]
	if len(conns) == 0 {
		var err error
		fmt.Printf("get conns from etcd remote, serviceName: %s\n", serviceName)
		addrs, err := s.GetConnsRemote(serviceName)
		if err != nil {
			return nil, err
		}
		if len(addrs) == 0 {
			return nil, fmt.Errorf("no conn for service %s, grpc server may not exist, local conn is %v, please check etcd server %v, key: %s", serviceName, s.localConns, s.etcdAddr, s.key)
		}
		for _, addr := range addrs {
			cc, err := grpc.DialContext(ctx, addr.Addr, append(s.options, opts...)...)
			if err != nil {
				fmt.Println("dialContext failed", err, "addr", addr.Addr, "opts", append(s.options, opts...))
				return nil, err
			}
			conns = append(conns, cc)
		}
		s.localConns[serviceName] = conns
	}
	return conns, nil
}
func (r *EtcdRegister) GetConn(ctx context.Context, serviceName string, opts ...grpc.DialOption) (grpc.ClientConnInterface, error) {
	newOpts := append(r.options, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, r.balancerName)))
	fmt.Printf("get conn from client, serviceName: %s\n", serviceName)
	return grpc.DialContext(ctx, fmt.Sprintf("%s:///%s", r.schema, serviceName), append(newOpts, opts...)...)
}
