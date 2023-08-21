package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpc-demo/demo-8/proto/friend"
	"net"
)

type Server struct {
}

func (s *Server) GetFriendInfo(ctx context.Context, info *friend.FriendBaseInfo) (*friend.RadarSearchPlayerInfo, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		requestID := md.Get("X-Request-Id")
		method := md.Get("method")
		pattern := md.Get("pattern")
		fmt.Println("method:", method, "pattern:", pattern, "requestID:", requestID)
	} else {
		return nil, status.Errorf(codes.DataLoss, "no metadata")
	}
	grpc.SetHeader(ctx, metadata.Pairs("code", "201"))
	return &friend.RadarSearchPlayerInfo{
		Distance:    6,
		PlayerId:    7,
		BubbleFrame: 8,
		Head:        9,
		HeadFrame:   10,
		NickName:    "你好",
	}, nil
}

func (s *Server) Test(ctx context.Context, info *friend.RadarSearchPlayerInfo) (*friend.RadarSearchPlayerInfo, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for k, v := range md {
			fmt.Println(k, v)
		}
		method := md.Get("method")
		pattern := md.Get("pattern")
		fmt.Println("method:", method, "pattern:", pattern)
	} else {
		return nil, status.Errorf(codes.DataLoss, "no metadata")
	}
	grpc.SetHeader(ctx, metadata.Pairs("code", "201"))
	return &friend.RadarSearchPlayerInfo{
		Distance:    1,
		PlayerId:    2,
		BubbleFrame: 3,
		Head:        4,
		HeadFrame:   5,
		NickName:    "你好",
	}, nil
}

func main() {
	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	friend.RegisterFriendServer(server, &Server{})
	listen, _ := net.Listen("tcp", ":8082")
	_ = server.Serve(listen)
}
