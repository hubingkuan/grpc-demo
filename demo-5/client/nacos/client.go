package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-demo/demo-5/config"
	"grpc-demo/demo-5/discoveryregisty/nacos"
	"grpc-demo/demo-5/interceptor"
	pb "grpc-demo/demo-5/proto"
	"log"
)

func main() {
	// round robin + 超时连接配置
	discoveryClient, err := nacos.NewClient(config.Config.Nacos.NamespaceID, config.Config.Nacos.Address, config.Config.Nacos.Schema, nacos.WithUserNameAndPassword(
		config.Config.Nacos.UserName,
		config.Config.Nacos.Password,
	), nacos.WithRoundRobin(), nacos.WithTimeout(5))
	if err != nil {
		log.Fatalln("init nacos client failed, err:", err)
	}

	// 客户端拦截器+ 不验证证书
	discoveryClient.AddOption(grpc.WithUnaryInterceptor(interceptor.RpcClientInterceptor), grpc.WithTransportCredentials(insecure.NewCredentials()))

	// 获取所有连接
	// conns, _ := r.GetConns(context.Background(), "helloServer")
	// fmt.Println(conns)

	for i := 0; i < 2; i++ {
		// 获取一个连接
		conn, err := discoveryClient.GetConn(context.Background(), "helloServer")
		if err != nil {
			panic(err)
		}
		client := pb.NewServerClient(conn)
		helloResponse, err := client.Hello(context.Background(), &pb.Empty{})
		if err != nil {
			fmt.Printf("err: %v", err)
			return
		}
		fmt.Println("resp: ", helloResponse)
	}
}
