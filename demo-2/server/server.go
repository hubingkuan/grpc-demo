package main

import (
	context "context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc-demo/demo-2/middleware/auth"
	"grpc-demo/demo-2/middleware/recovery"
	"grpc-demo/demo-2/middleware/validate"
	"grpc-demo/demo-2/proto/hello"
	"net"
	"os"
	"time"
)

type Server struct {
}

var _ hello.GreeterServer = (*Server)(nil)

func (s *Server) SayHello(ctx context.Context, person *hello.Person) (*hello.Person, error) {
	time.Sleep(2 * time.Second)
	// 发送header 和 trailer
	defer func() {
		trailer := metadata.Pairs("timestamp", time.Now().Format(time.StampNano))
		grpc.SetTrailer(ctx, trailer)
	}()

	header := metadata.New(map[string]string{"location": "MTV", "timestamp": time.Now().Format(time.StampNano)})
	grpc.SendHeader(ctx, header)

	return &hello.Person{
		Id: 999,
	}, nil
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	var opts []grpc.ServerOption
	//  设置一个拦截器 基于token校验和proto的validate校验
	opts = append(opts, grpc.ChainUnaryInterceptor(recovery.UnaryServerInterceptor(), auth.UnaryServerInterceptor(), validate.UnaryServerInterceptor()))
	opts = append(opts, grpc.ChainStreamInterceptor(recovery.StreamServerInterceptor(), auth.StreamServerInterceptor(), validate.StreamServerInterceptor()))

	server := grpc.NewServer(opts...)
	hello.RegisterGreeterServer(server, &Server{})
	listen, _ := net.Listen("tcp", ":8081")
	_ = server.Serve(listen)
}
