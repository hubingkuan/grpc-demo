package loadbalance

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
)

func Demo() {
	var options []grpc.DialOption
	options = append(options,
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		// 选项跳过了对服务器证书的验证
		grpc.WithTransportCredentials(insecure.NewCredentials()))
}
