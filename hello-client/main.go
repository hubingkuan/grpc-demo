/*package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	pb "grpc-demo/hello-server/proto"
)

type ClientTokenAuth struct {

}

// 获取元数据信息 context用于控制超时时间和取消 url是请求入口处的url
func (c ClientTokenAuth)GetRequestMetadata(ctx context.Context,url ... string)(map[string]string,error){
	return map[string]string{
		"appID":"maguahu",
		"appKey":"123",
	},nil
}

// 是否需要基于TLS认证进行安全传输
func (c ClientTokenAuth)RequireTransportSecurity()bool{
	return false
}


func main() {
	var opts []grpc.DialOption
	opts=append(opts,grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts=append(opts,grpc.WithPerRPCCredentials(new(ClientTokenAuth)))


	// 连接到server端
	conn, err := grpc.Dial("127.0.0.1:9090", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 建立连接
	client := pb.NewSayHelloClient(conn)
	// 这里创建metadata给服务端(所有的key都会转换成小写)
	md := metadata.Pairs("appid", "maguahu", "appkey", "12345")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	// 执行rpc调用
	response, err := client.SayHello(ctx, &pb.HelloRequest{RequestName: "maguahu"})
	fmt.Println(response.GetResponseMsg())
}*/