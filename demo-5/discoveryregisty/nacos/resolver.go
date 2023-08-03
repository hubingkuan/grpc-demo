package nacos

import "google.golang.org/grpc/resolver"

type Resolver struct {
	target resolver.Target
	cc     resolver.ClientConn
	addrs  []resolver.Address

	getConnsRemote func(serviceName string) (conns []resolver.Address, err error)
}

func (n NacosRegister) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	// TODO implement me
	panic("implement me")
}

func (r Resolver) ResolveNow(options resolver.ResolveNowOptions) {
}

func (r Resolver) Close() {
}
