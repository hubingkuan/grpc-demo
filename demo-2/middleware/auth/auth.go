package auth

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("method:", info.FullMethod)
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

func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		fmt.Println("method:", info.FullMethod)
		// 接受客户端的metadata
		md, ok := metadata.FromIncomingContext(stream.Context())
		if !ok {
			return status.Error(codes.DataLoss, "failed to get metadata")
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
			return status.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
		}
		return handler(srv, stream)
	}
}
