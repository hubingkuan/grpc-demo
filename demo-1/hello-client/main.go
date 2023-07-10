package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"grpc-demo/demo-1/hello-client/proto"
	pb "grpc-demo/demo-1/hello-server/proto"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	// 连接到server端
	conn, err := grpc.Dial("127.0.0.1:9091", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 建立连接
	client := pb.NewSayHelloClient(conn)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "version", "v1")
	// 这里创建metadata给服务端(所有的key都会转换成小写)
	md := metadata.Pairs("appid", "maguahu", "appkey", "12345")
	ctx = metadata.NewOutgoingContext(ctx, md)
	// 执行rpc调用
	response, err := client.SayHello2(ctx, &demoClient.HelloRequest{RequestName: "maguahu"})
	fmt.Println(response.GetResponseMsg())
}
