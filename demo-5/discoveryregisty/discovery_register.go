package discoveryregisty

import (
	"context"
	"google.golang.org/grpc"
)

type Conn interface {
	GetConns(ctx context.Context, serviceName string, opts ...grpc.DialOption) ([]grpc.ClientConnInterface, error)
	GetConn(ctx context.Context, serviceName string, opts ...grpc.DialOption) (grpc.ClientConnInterface, error)
	AddOption(opts ...grpc.DialOption)
	// do not use this method for call rpc
	GetClientLocalConns() map[string][]grpc.ClientConnInterface
}

type SvcDiscoveryRegistry interface {
	Conn
	// 注册服务
	Register(serviceName, host string, port int) error
	// 注销服务
	UnRegister() error
	// 注册配置
	RegisterConf2Registry(key string, conf []byte) error
	// 从注册中心获取配置
	GetConfFromRegistry(key string) ([]byte, error)
}
