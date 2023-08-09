package main

import (
	"context"
	"flag"
	"fmt"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"grpc-demo/demo-5/config"
	"grpc-demo/demo-5/discoveryregisty/etcd"
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

	// 拦截器+设置最大发送接受数据大小  默认接受数据大小为4M 发送数据大小为int32最大值
	var grpcOpts []grpc.ServerOption
	grpcOpts = append(grpcOpts, grpc.ChainUnaryInterceptor(interceptor.RpcServerInterceptor), grpc.MaxRecvMsgSize(1024*1024*10), grpc.MaxSendMsgSize(1024*1024*10))
	srv := grpc.NewServer(grpcOpts...)
	// 服务注册grpc服务器
	pb.RegisterServerServer(srv, Server{})
	r, err := etcd.NewClient(config.Config.Etcd.Address, config.Config.Etcd.Schema, etcd.WithUserNameAndPassword(
		config.Config.Etcd.UserName,
		config.Config.Etcd.Password,
	), etcd.WithTimeout(5))
	if err != nil {
		log.Fatalln("init etcd client failed, err:", err)
	}
	// 服务注册etcd
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
