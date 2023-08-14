package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-demo/demo-5/config"
	"grpc-demo/demo-5/discoveryregisty/zookeeper"
	"grpc-demo/demo-5/interceptor"
	pb "grpc-demo/demo-5/proto"
	"log"
)

func main() {
	discoveryClient, err := zookeeper.NewClient(config.Config.Zookeeper.Address, config.Config.Zookeeper.Schema, zookeeper.WithUserNameAndPassword(
		config.Config.Zookeeper.UserName,
		config.Config.Zookeeper.Password,
	), zookeeper.WithTimeout(5), zookeeper.WithRoundRobin())
	if err != nil {
		log.Fatalln("init zookeeper client failed, err:", err)
	}

	// 客户端拦截器+ 不验证证书
	discoveryClient.AddOption(grpc.WithUnaryInterceptor(interceptor.RpcClientInterceptor), grpc.WithTransportCredentials(insecure.NewCredentials()))

	// for i := 0; i < 10; i++ {
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
	// time.Sleep(time.Second * 40)
	// }
}
