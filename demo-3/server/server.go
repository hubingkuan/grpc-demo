package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	router "grpc-demo/demo-3/proto"
	"io"
	"log"
	"math"
	"net"
	"time"
)

type routerServer struct {
	router.UnimplementedRouteGuideServer
}

// 简单的rpc调用
func (r *routerServer) GetFeature(ctx context.Context, point *router.Point) (*router.Feature, error) {
	return &router.Feature{Location: point}, nil
}

// 流式返回数据
func (r *routerServer) ListFeatures(rectangle *router.Rectangle, serverStream router.RouteGuide_ListFeaturesServer) error {
	for i := 0; i < 10; i++ {
		if err := serverStream.Send(&router.Feature{
			Name:     "test",
			Location: &router.Point{Latitude: 1, Longitude: 2},
		}); err != nil {
			return err
		}
	}
	return nil
}

func calcDistance(p1 *router.Point, p2 *router.Point) int32 {
	const CordFactor float64 = 1e7
	const R = float64(6371000) // earth radius in metres
	lat1 := toRadians(float64(p1.Latitude) / CordFactor)
	lat2 := toRadians(float64(p2.Latitude) / CordFactor)
	lng1 := toRadians(float64(p1.Longitude) / CordFactor)
	lng2 := toRadians(float64(p2.Longitude) / CordFactor)
	dlat := lat2 - lat1
	dlng := lng2 - lng1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return int32(distance)
}
func toRadians(num float64) float64 {
	return num * math.Pi / float64(180)
}

// 流式接受客户端数据
func (r *routerServer) RecordRoute(clientStream router.RouteGuide_RecordRouteServer) error {
	var pointCount, featureCount, distance int32
	var lastPoint *router.Point
	startTime := time.Now()
	// 接受客户端传递的metadata
	md, ok := metadata.FromIncomingContext(clientStream.Context())
	if !ok {
		return status.Error(codes.DataLoss, "failed to get metadata")
	}
	fmt.Println(md)
	for {
		point, err := clientStream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			// 发送数据后关闭流
			return clientStream.SendAndClose(&router.RouteSummary{
				PointCount:   pointCount,
				FeatureCount: featureCount,
				Distance:     distance,
				ElapsedTime:  int32(endTime.Sub(startTime).Seconds()),
			})
		}
		if err != nil {
			return err
		}
		pointCount++
		if lastPoint != nil {
			distance += calcDistance(lastPoint, point)
		}
		lastPoint = point
	}
}

// 双向流式数据处理
func (r *routerServer) RouteChat(stream router.RouteGuide_RouteChatServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if err := stream.Send(in); err != nil {
			return err
		}
	}
}

func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return status.Errorf(codes.InvalidArgument, "no incoming metadata in rpc context")
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
	return handler(srv, ss)
}

func tokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
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

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	// 综合使用流式拦截器和普通拦截器
	opts = append(opts, grpc.UnaryInterceptor(tokenInterceptor), grpc.StreamInterceptor(streamInterceptor))

	// 如果需要多个流式拦截器使用 grpc.ChainStreamInterceptor()
	// 如果proto中既有
	grpcServer := grpc.NewServer(opts...)
	router.RegisterRouteGuideServer(grpcServer, &routerServer{})
	grpcServer.Serve(lis)
}
