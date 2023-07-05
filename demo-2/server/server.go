package main

import (
	context "context"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpc-demo/demo-2/proto/hello"
	"net"
	"os"
	"runtime/debug"
	"time"
)

type Server struct {
}

// 这里proto里面定义的包含了validate的message都已经实现了该接口
type Validator interface {
	Validate() error
}

var _ hello.GreeterServer = (*Server)(nil)

func (s *Server) SayHello(ctx context.Context, person *hello.Person) (*hello.Person, error) {
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
	opts = append(opts, grpc.ChainUnaryInterceptor(logInterceptor(), tokenInterceptor(), validateInterceptor()))

	server := grpc.NewServer(opts...)
	hello.RegisterGreeterServer(server, &Server{})
	listen, _ := net.Listen("tcp", ":8081")
	_ = server.Serve(listen)
}

func logInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		st := time.Now()
		defer func() {
			if recoverError := recover(); recoverError != nil {
				log.Error().Str("method", info.FullMethod).Interface("recover", recoverError).Bytes("stack", debug.Stack()).Interface("req", req).Interface("resp", resp).Err(err).TimeDiff("runtime", time.Now(), st).Send()
			} else {
				log.Info().Str("method", info.FullMethod).Interface("req", req).Interface("resp", resp).Err(err).TimeDiff("runtime", time.Now(), st).Send()
			}
		}()
		return handler(ctx, req)
	}
}

func validateInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
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
}

func tokenInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 获取调用的方法 /Greeter/SayHello
		method, ok := grpc.Method(ctx)
		if !ok {
			return nil, errors.New("missing method in incoming context")
		}
		fmt.Println("method:", method)
		// 接受客户端的metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.DataLoss, "failed to get metadata")
		}
		var (
			appid  string
			appkey string
		)
		if val, ok := md["appid"]; ok {
			appid = val[0]
		}
		if val, ok := md["appkey"]; ok {
			appkey = val[0]
		}
		if appid != "101010" || appkey != "i am key" {
			return nil, status.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
		}
		return handler(ctx, req)
	}
}
