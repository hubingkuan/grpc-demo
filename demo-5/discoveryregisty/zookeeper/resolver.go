package zookeeper

import (
	"context"
	"fmt"
	"google.golang.org/grpc/resolver"
	"strings"
)

type Resolver struct {
	target         resolver.Target
	cc             resolver.ClientConn
	addrs          []resolver.Address
	getConnsRemote func(serviceName string) (conns []resolver.Address, err error)
	watchZkChange  func(serviceName string)
}

func (r *Resolver) ResolveNowZK(o resolver.ResolveNowOptions) error {
	serviceName := strings.TrimLeft(r.target.URL.Path, "/")
	fmt.Println("start resolve now", "target", r.target, "serviceName", serviceName)
	newConns, err := r.getConnsRemote(serviceName)
	if err != nil {
		fmt.Println(context.Background(), "resolve now error", err, "target", r.target)
		return err
	}
	r.addrs = newConns
	if err = r.cc.UpdateState(resolver.State{Addresses: newConns}); err != nil {
		fmt.Println("UpdateState error, conns is nil from svr", err, "conns", newConns, "zk path", r.target.URL.Path)
		return err
	}
	go r.watchZkChange(serviceName)
	fmt.Println(context.Background(), "resolve now finished", "target", r.target, "conns", r.addrs)
	return nil
}

func (s *ZkClient) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	fmt.Printf("build resolver: %+v\n", target)
	r := &Resolver{}
	r.target = target
	r.cc = cc
	r.getConnsRemote = s.GetConnsRemote
	r.watchZkChange = s.watch
	err := r.ResolveNowZK(resolver.ResolveNowOptions{})
	s.lock.Lock()
	defer s.lock.Unlock()
	serviceName := strings.TrimLeft(target.URL.Path, "/")
	s.resolvers[serviceName] = r
	fmt.Printf("build resolver finished: %+v,  key: %s\n", target, serviceName)
	return r, err
}

func (r *Resolver) ResolveNow(o resolver.ResolveNowOptions) {}

func (s *Resolver) Close() {}
