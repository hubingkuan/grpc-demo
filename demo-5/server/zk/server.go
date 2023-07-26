package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpc-demo/demo-5/config"
	"grpc-demo/demo-5/discoveryregisty/zookeeper"
	"grpc-demo/demo-5/interceptor"
	pb "grpc-demo/demo-5/proto"
	"log"
	"net"
	"os"
	"os/signal"
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
	addr := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen(
		"tcp",
		addr,
	)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	var grpcOpts []grpc.ServerOption
	grpcOpts = append(grpcOpts, grpc.ChainUnaryInterceptor(interceptor.RpcServerInterceptor))
	srv := grpc.NewServer(grpcOpts...)
	// 服务注册grpc服务器
	pb.RegisterServerServer(srv, Server{})
	r, err := zookeeper.NewClient(config.Config.Zookeeper.Address, config.Config.Zookeeper.Schema, zookeeper.WithUserNameAndPassword(
		config.Config.Zookeeper.UserName,
		config.Config.Zookeeper.Password,
	), zookeeper.WithTimeout(5))
	if err != nil {
		log.Fatalln("init etcd client failed, err:", err)
	}
	// 服务注册zookeeper
	err = r.Register("helloServer", "127.0.0.1", port)
	if err != nil {
		log.Fatalln("register server failed, err:", err)
	}
	srv.Serve(listener)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	go func() {
		select {
		case <-signals:
			r.UnRegister()
		}
	}()

}
