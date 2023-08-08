package nacos

import (
	"context"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"net"
)

func (n *NacosRegister) watch(serviceName string) {
	n.namingClient.Subscribe(&vo.SubscribeParam{
		ServiceName: serviceName,
		Clusters:    []string{n.clusterName},
		GroupName:   n.groupName,
		SubscribeCallback: func(services []model.Instance, err error) {
			fmt.Println("subscribe callback")
			n.lock.Lock()
			defer n.lock.Unlock()
			n.flushResolverAndDeleteLocal(serviceName)
		},
	})
}

func (n *NacosRegister) GetConnsRemote(serviceName string) (conns []resolver.Address, err error) {
	instances, err := n.namingClient.SelectInstances(vo.SelectInstancesParam{
		Clusters:    []string{n.clusterName},
		ServiceName: serviceName,
		GroupName:   n.groupName,
		HealthyOnly: true,
	})
	if err != nil {
		return nil, err
	}
	for _, instance := range instances {
		addr := net.JoinHostPort(instance.Ip, fmt.Sprintf("%d", instance.Port))
		fmt.Println("get conns from remote", "conn:", addr)
		conns = append(conns, resolver.Address{Addr: addr, ServerName: serviceName})
	}
	return conns, nil
}

func (n *NacosRegister) GetConns(ctx context.Context, serviceName string, opts ...grpc.DialOption) ([]grpc.ClientConnInterface, error) {
	fmt.Printf("get conns from client, serviceName: %s\n", serviceName)
	n.lock.Lock()
	defer n.lock.Unlock()
	opts = append(n.options, opts...)
	conns := n.localConns[serviceName]
	if len(conns) == 0 {
		var err error
		fmt.Printf("get conns from etcd remote, serviceName: %s\n", serviceName)
		addrs, err := n.GetConnsRemote(serviceName)
		if err != nil {
			return nil, err
		}
		if len(addrs) == 0 {
			return nil, fmt.Errorf("no conn for service %s, grpc server may not exist, local conn is %v, please check nacos server %v, key: %s", serviceName, n.localConns, n.nacosAddr, n.key)
		}
		for _, addr := range addrs {
			cc, err := grpc.DialContext(ctx, addr.Addr, append(n.options, opts...)...)
			if err != nil {
				fmt.Println("dialContext failed", err, "addr", addr.Addr, "opts", append(n.options, opts...))
				return nil, err
			}
			conns = append(conns, cc)
		}
		n.localConns[serviceName] = conns
	}
	return conns, nil
}

func (n *NacosRegister) GetConn(ctx context.Context, serviceName string, opts ...grpc.DialOption) (grpc.ClientConnInterface, error) {
	newOpts := append(n.options, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, n.balancerName)))
	fmt.Printf("get conn from client, serviceName: %s\n", serviceName)
	return grpc.DialContext(ctx, fmt.Sprintf("%s:///%s", n.schema, serviceName), append(newOpts, opts...)...)
}
