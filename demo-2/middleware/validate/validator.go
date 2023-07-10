package validate

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 这里proto里面定义的包含了validate的message都已经实现了该接口
type Validator interface {
	Validate() error
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
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

func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		wrapper := &recvWrapper{
			ServerStream: stream,
		}
		return handler(srv, wrapper)
	}
}

// 将流包装  重写RecvMsg方法
type recvWrapper struct {
	grpc.ServerStream
}

func (s *recvWrapper) RecvMsg(m any) error {
	if err := s.ServerStream.RecvMsg(m); err != nil {
		return err
	}
	if p, ok := m.(Validator); ok {
		if err := p.Validate(); err != nil {
			return status.Error(codes.InvalidArgument, err.Error())
		}
	}
	return nil
}
