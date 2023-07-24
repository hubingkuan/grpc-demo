package etcd

import (
	"fmt"
	"google.golang.org/grpc/resolver"
	"strings"
)

type Resolver struct {
	target resolver.Target
	cc     resolver.ClientConn
	addrs  []resolver.Address

	getConnsRemote  func(serviceName string) (conns []resolver.Address, err error)
	watchEtcdChange func(serviceName string, addrList []resolver.Address)
}

func (r *Resolver) ResolveNowEtcd(o resolver.ResolveNowOptions) error {
	serviceName := strings.TrimLeft(r.target.URL.Path, "/")
	fmt.Printf("start resolve now, target:%v ,serviceName:%s \n", r.target, serviceName)
	newConns, err := r.getConnsRemote(serviceName)
	if err != nil {
		fmt.Println("resolve now error", err, "target", r.target)
		return err
	}
	r.addrs = newConns
	if err = r.cc.UpdateState(resolver.State{Addresses: newConns}); err != nil {
		fmt.Println("UpdateState error, conns is nil from svr:", err, " conns:", newConns, " server path:", r.target.URL.Path)
		return err
	}
	go r.watchEtcdChange(serviceName, newConns)
	fmt.Println("resolve now finished", "target", r.target, "conns", r.addrs)
	return nil
}

func (s *EtcdRegister) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	if s.cli == nil {
		return nil, fmt.Errorf("etcd clientv3 client failed, etcd:%s", target)
	}
	fmt.Printf("build resolver: %+v\n", target)
	r := &Resolver{}
	r.target = target
	r.cc = cc
	r.getConnsRemote = s.GetConnsRemote
	r.watchEtcdChange = s.watch
	err := r.ResolveNowEtcd(resolver.ResolveNowOptions{})
	s.lock.Lock()
	defer s.lock.Unlock()
	serviceName := strings.TrimLeft(target.URL.Path, "/")
	s.resolvers[serviceName] = r
	fmt.Printf("build resolver finished: %+v,key: %s\n", target, serviceName)
	return r, err
}

func (r *Resolver) ResolveNow(o resolver.ResolveNowOptions) {}

func (s *Resolver) Close() {}
