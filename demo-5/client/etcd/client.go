package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-demo/demo-5/config"
	"grpc-demo/demo-5/discoveryregisty/etcd"
	"grpc-demo/demo-5/interceptor"
	pb "grpc-demo/demo-5/proto"
	"log"
	"time"
)

func main() {
	// round robin + 超时连接配置
	discoveryClient, err := etcd.NewClient(config.Config.Etcd.Address, config.Config.Etcd.Schema, etcd.WithUserNameAndPassword(
		config.Config.Etcd.UserName,
		config.Config.Etcd.Password,
	), etcd.WithRoundRobin(), etcd.WithTimeout(5))
	if err != nil {
		log.Fatalln("init etcd client failed, err:", err)
	}

	// 客户端拦截器+ 不验证证书 设置最大发送接受数据大小
	discoveryClient.AddOption(grpc.WithUnaryInterceptor(interceptor.RpcClientInterceptor), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*10), grpc.MaxCallSendMsgSize(1024*1024*10)))

	// 获取所有连接
	// conns, _ := r.GetConns(context.Background(), "helloServer")
	// fmt.Println(conns)

	for i := 0; i < 100; i++ {
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
		time.Sleep(30 * time.Second)
	}
}
