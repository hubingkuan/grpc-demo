package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc-demo/demo-7/proto"
	requestid "grpc-demo/demo-7/requestId"
	"grpc-demo/demo-7/tracing"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	tracer, closer := tracing.Init("demo-7-client")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			func(
				ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
			) (resp interface{}, err error) {
				ctx = requestid.NewWithContext(ctx, "")
				resp, err = handler(ctx, req)
				return resp, err
			},
		),
	)

	conn, err := grpc.Dial(
		"localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			func(
				ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
				invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
			) error {
				id := requestid.RequestID(ctx)
				ctx = metadata.AppendToOutgoingContext(ctx, "x-request-id", id)
				return invoker(ctx, method, req, reply, cc, opts...)
			}, grpc_opentracing.UnaryClientInterceptor(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewServerClient(conn)
	hello, err := client.Hello(context.Background(), &pb.Empty{})
	if err != nil {
		log.Println(err.Error())
	} else {
		fmt.Println("resp:", hello)
	}
}
