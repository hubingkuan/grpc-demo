package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"net"
	"strings"
)

func RpcServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	fmt.Printf("req:%s , call %s\n", req, info.FullMethod)
	ip, err := getClietIP(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("client ip: ", ip)
	/*	md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.New(codes.InvalidArgument, "missing metadata").Err()
		}
		// 必要参数解析
		if opts := md.Get("OperationID"); len(opts) != 1 || opts[0] == "" {
			return nil, status.New(codes.InvalidArgument, "operationID error").Err()
		} else {
			ctx = context.WithValue(ctx, "OperationID", opts[0])
		}
		// 非必要参数解析
		if opts := md.Get("OpUserPlatform"); len(opts) == 1 {
			ctx = context.WithValue(ctx, "OpUserPlatform", opts[0])
		}
		if opts := md.Get("ConnID"); len(opts) == 1 {
			ctx = context.WithValue(ctx, "ConnID", opts[0])
		}*/
	resp, err = handler(ctx, req)
	return resp, err
}

func getClietIP(ctx context.Context) (string, error) {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		return "", fmt.Errorf("[getClinetIP] invoke FromContext() failed")
	}
	if pr.Addr == net.Addr(nil) {
		return "", fmt.Errorf("[getClientIP] peer.Addr is nil")
	}
	addSlice := strings.Split(pr.Addr.String(), ":")
	return addSlice[0], nil
}
