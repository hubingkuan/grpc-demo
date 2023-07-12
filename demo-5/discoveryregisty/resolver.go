package discoveryregisty

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

func (r *Resolver) SetGetConnsRemote(getConnsRemote func(serviceName string) (conns []resolver.Address, err error)) {
	r.getConnsRemote = getConnsRemote
}

func (r *Resolver) SetTarget(arget resolver.Target) {
	r.target = arget
}

func (r *Resolver) SetCc(cc resolver.ClientConn) {
	r.cc = cc
}

func (r *Resolver) SetAddrs(addrs []resolver.Address) {
	r.addrs = addrs
}

func (r *Resolver) ResolveNowZK(o resolver.ResolveNowOptions) {
	fmt.Println("start resolve now", "target", r.target, "cc", r.cc.UpdateState, "serviceName", strings.TrimLeft(r.target.URL.Path, "/"))
	newConns, err := r.getConnsRemote(strings.TrimLeft(r.target.URL.Path, "/"))
	if err != nil {
		fmt.Println("resolve now error", err, "target", r.target)
		return
	}
	r.addrs = newConns
	if err := r.cc.UpdateState(resolver.State{Addresses: newConns}); err != nil {
		fmt.Println("UpdateState error, conns is nil from svr", err, "conns", newConns, "zk path", r.target.URL.Path)
		return
	}
	fmt.Println("resolve now finished", "target", r.target, "conns", r.addrs)
}

func (r *Resolver) ResolveNow(o resolver.ResolveNowOptions) {}

func (s *Resolver) Close() {}
