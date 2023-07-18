package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-demo/demo-4/proto"
	"grpc-demo/demo-5/config"
	"grpc-demo/demo-5/discoveryregisty/etcd"
	"net"
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
	listener, err := net.Listen(
		"tcp",
		net.JoinHostPort("127.0.0.1", "8081"),
	)
	if err != nil {
		panic(err)
	}
	var grpcOpts []grpc.ServerOption
	grpcOpts = append(grpcOpts, grpc.ChainUnaryInterceptor(UnaryInterceptor()))
	srv := grpc.NewServer(grpcOpts...)
	pb.RegisterServerServer(srv, Server{})
	r, _ := etcd.NewClient(config.Config.Etcd.EtcdAddr, config.Config.Etcd.EtcdSchema)
	r.Register("helloServer", "127.0.0.1", 8081)
	srv.Serve(listener)
}

func UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Printf("call %s\n", info.FullMethod)
		resp, err = handler(ctx, req)
		return resp, err
	}
}
