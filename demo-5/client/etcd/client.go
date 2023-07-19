package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-demo/demo-4/proto"
	"grpc-demo/demo-5/config"
	"grpc-demo/demo-5/discoveryregisty/etcd"
)

func main() {
	// round robin + 默认中间件
	r, _ := etcd.NewClient(config.Config.Etcd.EtcdAddr, config.Config.Etcd.EtcdSchema, etcd.WithRoundRobin(), etcd.WithOptions(grpc.WithUnaryInterceptor(RpcClientInterceptor)))
	r.AddOption(grpc.WithTransportCredentials(insecure.NewCredentials()))
	// 获取所有连接
	// conns, _ := r.GetConns(context.Background(), "helloServer")
	// fmt.Println(conns)

	for i := 0; i < 10; i++ {
		// 获取一个连接
		conn, err := r.GetConn(context.Background(), "helloServer")
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

func RpcClientInterceptor(
	ctx context.Context,
	method string,
	req, resp interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) (err error) {
	if ctx == nil {
		return errors.New("call rpc request context is nil")
	}
	// ctx, err = getRpcContext(ctx, method)  包装ctx传值 metadata.NewOutgoingContext
	if err != nil {
		return err
	}
	fmt.Println(ctx, "get rpc ctx success", "conn target", cc.Target())
	err = invoker(ctx, method, req, resp, cc, opts...)
	if err == nil {
		fmt.Println(ctx, "rpc client resp", "funcName", method, "resp", rpcString(resp))
		return nil
	}
	fmt.Println(ctx, "rpc resp error", err)
	return err
}

func rpcString(v interface{}) string {
	if s, ok := v.(interface{ String() string }); ok {
		return s.String()
	}
	return fmt.Sprintf("%+v", v)
}
