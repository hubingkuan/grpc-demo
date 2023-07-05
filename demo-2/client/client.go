package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"grpc-demo/demo-2/proto/hello"
	"log"
	"time"
)

type ClientTokenAuth struct {
}

// 获取元数组信息，也就是客户端提供的key，value对，context用于控制超时和取消，uri是请求入口处的uri，
func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, url ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "101010",
		"appKey": "i am key",
	}, nil
}

// 是否需要基于TLS认证进行安全传输
func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	conn, _ := grpc.Dial("localhost:8081", opts...)
	defer conn.Close()
	client := hello.NewGreeterClient(conn)

	// 发送metadata
	md := metadata.Pairs("timestamp", time.Now().Format(time.StampNano))
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// 定义接收header和trailer
	var header, trailer metadata.MD
	res, err := client.SayHello(ctx, &hello.Person{
		// 这里故意传递一个不符合规范的值
		Id:    1000,
		Email: "fanqiechaodan@fanqiechaodan.com",
		Name:  "番茄炒蛋",
		Home: &hello.Person_Location{
			Lat: 23,
			Lng: 45,
		},
	}, grpc.Header(&header), grpc.Trailer(&trailer))

	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(res.Id)

	if t, ok := header["timestamp"]; ok {
		fmt.Printf("timestamp from header:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Fatal("timestamp expected but doesn't exist in header")
	}

	if t, ok := trailer["timestamp"]; ok {
		fmt.Printf("timestamp from trailer:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Fatal("timestamp expected but doesn't exist in trailer")
	}
}
