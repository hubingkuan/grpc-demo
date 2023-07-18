package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"io"
	"strings"
)

func (r *EtcdRegister) watch(serviceName string, addrList []resolver.Address) {
	rch := r.cli.Watch(context.Background(), GetPrefix(r.schema, serviceName), clientv3.WithPrefix())
	for n := range rch {
		flag := 0
		for _, ev := range n.Events {
			switch ev.Type {
			case mvccpb.PUT:
				if !exists(addrList, string(ev.Kv.Value)) {
					flag = 1
					addrList = append(addrList, resolver.Address{Addr: string(ev.Kv.Value)})
					fmt.Println("after add, new list: ", addrList)
				}
			case mvccpb.DELETE:
				fmt.Println("remove addr key: ", string(ev.Kv.Key), "value:", string(ev.Kv.Value))
				i := strings.LastIndexAny(string(ev.Kv.Key), "/")
				if i < 0 {
					return
				}
				t := string(ev.Kv.Key)[i+1:]
				fmt.Println("remove addr key: ", string(ev.Kv.Key), "value:", string(ev.Kv.Value), "addr:", t)
				if s, ok := remove(addrList, t); ok {
					flag = 1
					addrList = s
					fmt.Println("after remove, new list: ", addrList)
				}
			}
		}

		if flag == 1 {
			r.lock.Lock()
			defer r.lock.Unlock()
			// 清空本地缓存数据
			r.localConns[serviceName] = r.localConns[serviceName][:0]
			r.resolvers[serviceName].cc.UpdateState(resolver.State{Addresses: addrList})
			r.resolvers[serviceName].addrs = addrList
			fmt.Println("update: ", addrList)
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

func (r *EtcdRegister) CloseConn(conn grpc.ClientConnInterface) {
	if closer, ok := conn.(io.Closer); ok {
		closer.Close()
	}
}
