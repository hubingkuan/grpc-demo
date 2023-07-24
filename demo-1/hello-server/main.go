package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpc-demo/demo-1/hello-client/proto"
	"net"
)

type server struct {
	demoClient.UnimplementedSayHelloServer
}

func (s *server) SayHello2(ctx context.Context, req *demoClient.HelloRequest) (*demoClient.HelloResponse, error) {
	fmt.Println("客户端参数:", req.RequestName)
	// 获取元数据信息   FromIncomingContext用于服务端获取request中的meta，FromOutgoingContext用于客户端获取自己即将发出的metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "未传递token")
	}
	var appId string
	var appKey string
	if v, ok := md["appid"]; ok {
		appId = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}
	if appId != "maguahu" && appKey != "123" {
		return nil, errors.New("token 不正确")
	}
	return &demoClient.HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

func main() {
	// 开启端口
	listen, err := net.Listen("tcp", ":9091")
	if err != nil {
		panic(err)
	}
	// 创建grpc服务 并且注册一个响应时间的中间件
	grpcServer := grpc.NewServer(
	/*grpc.ChainUnaryInterceptor(
	func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("one before")
		t1 := time.Now()
		resp, err = handler(ctx, req)
		fmt.Println("耗时:", time.Since(t1))
		fmt.Println("one after")
		return resp, err
	}, func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("two before")
		t1 := time.Now()
		resp, err = handler(ctx, req)
		fmt.Println("耗时:", time.Since(t1))
		fmt.Println("two after")
		return resp, err
	})*/)
	// 注册服务
	demoClient.RegisterSayHelloServer(grpcServer, &server{})
	// 启动服务
	grpcServer.Serve(listen)
}
