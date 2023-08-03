package nacos

import (
	"context"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"google.golang.org/grpc"
)

func (n NacosRegister) GetConns(ctx context.Context, serviceName string, opts ...grpc.DialOption) ([]grpc.ClientConnInterface, error) {
	fmt.Printf("get conns from client, serviceName: %s\n", serviceName)
	service, err := n.namingClient.GetService(vo.GetServiceParam{
		Clusters:    nil,
		ServiceName: "",
		GroupName:   "",
	})
	if err != nil {
		fmt.Println()
	}

	// TODO implement me
	panic("implement me")
}

func (n NacosRegister) GetConn(ctx context.Context, serviceName string, opts ...grpc.DialOption) (grpc.ClientConnInterface, error) {
	// TODO implement me
	panic("implement me")
}
