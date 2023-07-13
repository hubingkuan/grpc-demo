package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-demo/demo-4/proto"
	"grpc-demo/demo-5/config"
	"grpc-demo/demo-5/discoveryregisty/etcd"
)

type Server struct {
}

func (s Server) Hello(ctx context.Context, request *pb.Empty) (*pb.HelloResponse, error) {
	resp := pb.HelloResponse{Hello: "hello client."}
	return &resp, nil
}

func (s Server) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	resp := pb.RegisterResponse{}
	resp.Uid = fmt.Sprintf("%s.%s", request.GetName(), request.GetPassword())
	return &resp, nil
}

func main() {
	var grpcOpts []grpc.ServerOption
	srv := grpc.NewServer(grpcOpts...)
	pb.RegisterServerServer(srv, Server{})
	r, _ := etcd.NewClient(config.Config.Etcd.EtcdAddr, config.Config.Etcd.EtcdSchema)
	r.Register("helloServer", "127.0.0.1", 8081)
}
