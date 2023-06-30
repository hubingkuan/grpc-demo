package main

import (
	context "context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-demo/demo-2/proto/hello"
	"net"
)

type Server struct {
}

// 这里proto里面定义的包含了validate的message都已经实现了该接口
type Validator interface {
	Validate() error
}

func (s Server) SayHello(ctx context.Context, person *hello.Person) (*hello.Person, error) {
	return &hello.Person{
		Id: 999,
	}, nil
}

func main() {
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 使用Validator不使用Person;是为了避免硬编码
		// 这样拦截器每个需要验证的req都可以使用
		if p, ok := req.(Validator); ok {
			if err := p.Validate(); err != nil {
				// 参数没有通过验证时;返回一个错误不继续向下执行
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}
		// 通过验证后继续向下执行
		return handler(ctx, req)
	}
	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor))
	hello.RegisterGreeterServer(server, &Server{})
	listen, _ := net.Listen("tcp", ":8081")
	_ = server.Serve(listen)
}
