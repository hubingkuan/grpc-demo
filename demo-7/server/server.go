package main

import (
	"context"
	"net"
	"time"

	pb "grpc-demo/demo-7/proto"
	requestid "grpc-demo/demo-7/requestId"
	"grpc-demo/demo-7/tracing"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type Server struct {
}

func (s *Server) Hello(ctx context.Context, empty *pb.Empty) (*pb.HelloResponse, error) {
	grpclog.Info(ctx)
	s.doSomething(ctx)
	grpclog.Info(ctx)
	return &pb.HelloResponse{
		Hello: "maguahu",
	}, nil
}

func (s *Server) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	grpclog.Info(ctx)
	s.doSomething(ctx)
	grpclog.Info(ctx)
	return &pb.RegisterResponse{
		Uid: "1111",
	}, nil
}

func (s *Server) doSomething(ctx context.Context) {
	sp := opentracing.SpanFromContext(ctx)
	if sp != nil {
		opentracing.StartSpan("doSomething")
	} else {
		spContext := sp.Context()
		opentracing.StartSpan("doSomething", opentracing.ChildOf(spContext))
	}
	defer sp.Finish()
	sp.SetTag("age", "18")
	time.Sleep(time.Second * 2)
}

func main() {
	tracer, closer := tracing.Init("demo-7-server")
	defer closer.Close()
	// otel设置全局收集器
	opentracing.SetGlobalTracer(tracer)

	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			func(
				ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
			) (resp interface{}, err error) {
				ctx = requestid.NewWithContext(ctx, "")
				resp, err = handler(ctx, req)
				return resp, err
			}, grpc_opentracing.UnaryServerInterceptor(),
		),
	)
	pb.RegisterServerServer(srv, &Server{})
	listen, _ := net.Listen("tcp", ":8081")
	_ = srv.Serve(listen)
}
