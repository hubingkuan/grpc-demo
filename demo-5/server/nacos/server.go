package main

import (
	"context"
	"flag"
	"fmt"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"grpc-demo/demo-5/config"
	"grpc-demo/demo-5/discoveryregisty/nacos"
	"grpc-demo/demo-5/interceptor"
	pb "grpc-demo/demo-5/proto"
	"log"
	"net"
	"net/http"
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
	r, err := nacos.NewClient(config.Config.Nacos.NamespaceID, config.Config.Nacos.Address, config.Config.Nacos.Schema, nacos.WithUserNameAndPassword(
		config.Config.Nacos.UserName,
		config.Config.Nacos.Password,
	), nacos.WithRoundRobin(), nacos.WithTimeout(5))
	if err != nil {
		log.Fatalln("init nacos client failed, err:", err)
	}
	// 服务注册nacos
	err = r.Register("helloServer", "127.0.0.1", port)
	if err != nil {
		log.Fatalln("register server failed, err:", err)
	}
	srv.Serve(listener)

	// 服务监控(grpc自带)
	go startTrace()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	go func() {
		select {
		case <-signals:
			r.UnRegister()
		}
	}()
}

func startTrace() {
	// localhost:50051/debug/events
	// localhost:50051/debug/requests
	fmt.Println("start trace")
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}
	go http.ListenAndServe(":50051", nil)
	fmt.Println("Trace listen on 50051")
}
