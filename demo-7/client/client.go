package main

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	pb "grpc-demo/demo-7/proto"
	requestid "grpc-demo/demo-7/requestId"
	"grpc-demo/demo-7/tracing"
	"log"
)

func main() {
	tracer, closer := tracing.Init("helloService")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	grpc.NewServer(grpc.ChainUnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx = requestid.NewWithContext(ctx, "")
		resp, err = handler(ctx, req)
		return resp, err
	}))

	client, err := grpc.Dial("localhost:50001", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			id := requestid.RequestID(ctx)
			ctx = metadata.AppendToOutgoingContext(ctx, "x-request-id", id)
			return invoker(ctx, method, req, reply, cc, opts...)
		}))
	if err != nil {
		log.Fatal(err)
	}

	helloClient := pb.NewServerClient(client)

}
