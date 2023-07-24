package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"grpc-demo/demo-5/config"
	"grpc-demo/demo-5/discoveryregisty/zookeeper"
	pb "grpc-demo/demo-5/proto"
	"log"
	"time"
)

func main() {
	client, err := zookeeper.NewClient(config.Config.Zookeeper.Address, config.Config.Zookeeper.Schema, zookeeper.WithUserNameAndPassword(
		config.Config.Zookeeper.UserName,
		config.Config.Zookeeper.Password,
	), zookeeper.WithTimeout(5), zookeeper.WithRoundRobin(), zookeeper.WithFreq(time.Hour))
	if err != nil {
		log.Fatalln("init zookeeper client failed, err:", err)
	}

	// 客户端拦截器+ 不验证证书
	client.AddOption(grpc.WithUnaryInterceptor(RpcClientInterceptor), grpc.WithTransportCredentials(insecure.NewCredentials()))

	for i := 0; i < 10; i++ {
		// 获取一个连接
		conn, err := client.GetConn(context.Background(), "helloServer")
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
	// 包装ctx传值 metadata.NewOutgoingContext
	ctx, err = getRpcContext(ctx, method)
	if err != nil {
		return err
	}
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

func getRpcContext(ctx context.Context, method string) (context.Context, error) {
	md := metadata.Pairs()
	/*	// 必要参数传值
		operationID, ok := ctx.Value("OperationID").(string)
		if !ok {
			fmt.Println(ctx, "ctx missing operationID", errors.New("ctx missing operationID"), "funcName", method)
			return nil, errors.New("ctx missing operationID")
		}
		md.Set("OperationID", operationID)
		// 非必要参数传值
		connID, ok := ctx.Value("ConnID").(string)
		if ok {
			md.Set("ConnID", connID)
		}*/
	return metadata.NewOutgoingContext(ctx, md), nil
}

func rpcString(v interface{}) string {
	if s, ok := v.(interface{ String() string }); ok {
		return s.String()
	}
	return fmt.Sprintf("%+v", v)
}
