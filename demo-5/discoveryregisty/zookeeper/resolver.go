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
}

func (r *Resolver) ResolveNowZK(o resolver.ResolveNowOptions) {
	fmt.Println(
		context.Background(),
		"start resolve now",
		"target",
		r.target,
		"serviceName",
		strings.TrimLeft(r.target.URL.Path, "/"),
	)
	newConns, err := r.getConnsRemote(strings.TrimLeft(r.target.URL.Path, "/"))
	if err != nil {
		fmt.Println(context.Background(), "resolve now error", err, "target", r.target)
		return
	}
	r.addrs = newConns
	if err := r.cc.UpdateState(resolver.State{Addresses: newConns}); err != nil {
		fmt.Println(
			context.Background(),
			"UpdateState error, conns is nil from svr",
			err,
			"conns",
			newConns,
			"zk path",
			r.target.URL.Path,
		)
		return
	}
	fmt.Println(context.Background(), "resolve now finished", "target", r.target, "conns", r.addrs)
}

func (s *ZkClient) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	fmt.Printf("build resolver: %+v\n", target)
	r := &Resolver{}
	r.target = target
	r.cc = cc
	r.getConnsRemote = s.GetConnsRemote
	r.ResolveNowZK(resolver.ResolveNowOptions{})
	s.lock.Lock()
	defer s.lock.Unlock()
	serviceName := strings.TrimLeft(target.URL.Path, "/")
	s.resolvers[serviceName] = r
	fmt.Printf("build resolver finished: %+v,  key: %s\n", target, serviceName)
	return r, nil
}

func (r *Resolver) ResolveNow(o resolver.ResolveNowOptions) {}

func (s *Resolver) Close() {}
