package nacos

import (
	"fmt"
	"google.golang.org/grpc/resolver"
	"strings"
)

type Resolver struct {
	target resolver.Target
	cc     resolver.ClientConn
	addrs  []resolver.Address

	getConnsRemote func(serviceName string) (conns []resolver.Address, err error)
}

func (r *Resolver) ResolveNowNacos(o resolver.ResolveNowOptions) {
	serviceName := strings.TrimLeft(r.target.URL.Path, "/")
	fmt.Printf("start resolve now, target:%v ,serviceName:%s \n", r.target, serviceName)
	newConns, err := r.getConnsRemote(serviceName)
	if err != nil {
		fmt.Println("resolve now error", err, "target", r.target)
		return
	}
	r.addrs = newConns
	if err = r.cc.UpdateState(resolver.State{Addresses: newConns}); err != nil {
		fmt.Println("UpdateState error, conns is nil from svr:", err, " conns:", newConns, " server path:", r.target.URL.Path)
		return
	}
	fmt.Println("resolve now finished", "target", r.target, "conns", r.addrs)
	return
}

func (n *NacosRegister) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	fmt.Printf("build resolver: %+v\n", target)
	r := &Resolver{}
	r.target = target
	r.cc = cc
	r.getConnsRemote = n.GetConnsRemote
	r.ResolveNowNacos(resolver.ResolveNowOptions{})
	n.lock.Lock()
	defer n.lock.Unlock()
	serviceName := strings.TrimLeft(target.URL.Path, "/")
	n.resolvers[serviceName] = r
	go n.watch(serviceName)
	fmt.Printf("build resolver finished: %+v,key: %s\n", target, serviceName)
	return r, nil
}

func (r *Resolver) ResolveNow(options resolver.ResolveNowOptions) {
}

func (r *Resolver) Close() {
}
