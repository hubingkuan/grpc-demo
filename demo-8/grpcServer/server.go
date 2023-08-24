package main

import (
	"context"
	"fmt"
	"net"

	"grpc-demo/demo-8/proto/friend"
	"grpc-demo/demo-8/proto/hello"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, person *hello.Person) (*hello.Person, error) {
	return &hello.Person{
		Id:    1,
		Email: "maguahu@qq.com",
		Name:  "maguahu",
		Home:  nil,
	}, nil
}

func (s *Server) GetFriendInfo(
	ctx context.Context, info *friend.FriendBaseInfo,
) (*friend.RadarSearchPlayerInfo, error) {
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
	listen, err := net.Listen("tcp", ":8083")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := listen.Close(); err != nil {
			glog.Errorf("Failed to close %s %s: %v", "tcp", ":8083", err)
		}
	}()
	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	friend.RegisterFriendServer(server, &Server{})
	hello.RegisterGreeterServer(server, &Server{})
	_ = server.Serve(listen)
}
