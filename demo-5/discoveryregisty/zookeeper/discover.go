package zookeeper

import (
	"context"
	"fmt"
	"github.com/go-zookeeper/zk"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

func (s *ZkClient) watch(serviceName string) {
	for {
		fmt.Println("watching....")
		_, _, events, err := s.conn.ChildrenW(s.getPath(serviceName))
		if err != nil {
			fmt.Println("children watch error", err)
			return
		}
		if event, ok := <-events; ok {
			switch event.Type {
			case zk.EventNodeChildrenChanged:
				fmt.Printf("zk event: %s\n", event.Path)
				s.lock.Lock()
				s.flushResolverAndDeleteLocal(serviceName)
				s.lock.Unlock()
				fmt.Printf("zk event handle success: %s\n", event.Path)
			case zk.EventNodeCreated:
			case zk.EventNodeDeleted:
			case zk.EventNotWatching:
			case zk.EventSession:
			}
		}
	}
}

func (s *ZkClient) GetConnsRemote(serviceName string) (conns []resolver.Address, err error) {
	path := s.getPath(serviceName)
	childNodes, _, err := s.conn.Children(path)
	if err != nil {
		return nil, errors.Wrap(err, "get children error")
	} else {
		for _, child := range childNodes {
			fullPath := path + "/" + child
			data, _, err := s.conn.Get(fullPath)
			if err != nil {
				if err == zk.ErrNoNode {
					return nil, errors.Wrap(err, "this is zk ErrNoNode")
				}
				return nil, errors.Wrap(err, "get children error")
			}
			fmt.Println("get conns from remote", "conn", string(data))
			conns = append(conns, resolver.Address{Addr: string(data), ServerName: serviceName})
		}
	}
	return conns, nil
}

func (s *ZkClient) GetConns(ctx context.Context, serviceName string, opts ...grpc.DialOption) ([]grpc.ClientConnInterface, error) {
	fmt.Printf("get conns from client, serviceName: %s\n", serviceName)
	s.lock.Lock()
	defer s.lock.Unlock()
	opts = append(s.options, opts...)
	conns := s.localConns[serviceName]
	if len(conns) == 0 {
		var err error
		fmt.Printf("get conns from zk remote, serviceName: %s\n", serviceName)
		addrs, err := s.GetConnsRemote(serviceName)
		if err != nil {
			s.lock.Unlock()
			return nil, err
		}
		if len(addrs) == 0 {
			return nil, fmt.Errorf("no conn for service %s, grpc server may not exist, local conn is %v, please check zookeeper server %v, path: %s", serviceName, s.localConns, s.zkServers, s.schema)
		}
		for _, addr := range addrs {
			cc, err := grpc.DialContext(ctx, addr.Addr, append(s.options, opts...)...)
			if err != nil {
				fmt.Println("dialContext failed", err, "addr", addr.Addr, "opts", append(s.options, opts...))
				return nil, errors.Wrap(err, fmt.Sprintf("conns dialContext error, conn: %s", addr.Addr))
			}
			conns = append(conns, cc)
		}
		s.localConns[serviceName] = conns
	}
	return conns, nil
}

func (s *ZkClient) GetConn(ctx context.Context, serviceName string, opts ...grpc.DialOption) (grpc.ClientConnInterface, error) {
	newOpts := append(s.options, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, s.balancerName)))
	fmt.Printf("get conn from client, serviceName: %s\n", serviceName)
	return grpc.DialContext(ctx, fmt.Sprintf("%s:///%s", s.schema, serviceName), append(newOpts, opts...)...)
}
