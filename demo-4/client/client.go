package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-demo/demo-4/proto"
	"log"
	"time"
)

const etcdUrl = "http://localhost:2379"
const serviceName = "chihuo/server"

func main() {
	etcdClient, err := clientv3.NewFromURL(etcdUrl)
	if err != nil {
		panic(err)
	}
	etcdResolver, err := resolver.NewBuilder(etcdClient)

	conn, err := grpc.Dial(fmt.Sprintf("etcd:///%s", serviceName), grpc.WithResolvers(etcdResolver), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	ServerClient := pb.NewServerClient(conn)

	for {
		helloResponse, err := ServerClient.Hello(context.Background(), &pb.Empty{})
		if err != nil {
			fmt.Printf("err: %v", err)
			return
		}

		log.Println(helloResponse, err)
		time.Sleep(500 * time.Millisecond)
	}
}
