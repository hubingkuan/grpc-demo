package main

import (
	context "context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpc-demo/demo-4/proto"
	"log"
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
	var port int
	flag.IntVar(&port, "port", 8001, "port")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("启动grpc server失败")
		return
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(UnaryInterceptor()))

	pb.RegisterServerServer(grpcServer, Server{})

	log.Printf("service start port %d\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("启动grpc server失败")
	}
}

func UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Printf("call %s\n", info.FullMethod)
		resp, err = handler(ctx, req)
		return resp, err
	}
}
