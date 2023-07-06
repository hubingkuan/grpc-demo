package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	pb "grpc-demo/demo-4/proto"
	"log"
	"time"
)

func main() {
	bd := &ChihuoBuilder{addrs: map[string][]string{"api": []string{"localhost:8001", "localhost:8002", "localhost:8003"}}}
	// 注册自定义resolver
	resolver.Register(bd)
	conn, err := grpc.Dial("chihuo://api", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))

	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	ServerClient := pb.NewServerClient(conn)

	for {
		helloRespone, err := ServerClient.Hello(context.Background(), &pb.Empty{})
		if err != nil {
			fmt.Printf("err: %v", err)
			return
		}

		log.Println(helloRespone, err)
		time.Sleep(500 * time.Millisecond)
	}
}
