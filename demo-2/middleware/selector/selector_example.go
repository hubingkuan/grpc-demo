package selector

import (
	"context"
	"google.golang.org/grpc"
	"grpc-demo/demo-2/middleware/ratelimit"
)

func healthSkip(_ context.Context, c CallMeta) bool {
	return c.FullMethod() != "/ping.v1.PingService/Health"
}

func Example_ratelimit() {
	limiter := &ratelimit.ChannelTokenLimiter{}
	_ = grpc.NewServer(
		// 如果请求路径是/ping.v1.PingService/Health 则走ratelimit的拦截器  否则直接调用
		grpc.ChainUnaryInterceptor(
			UnaryServerInterceptor(ratelimit.UnaryServerInterceptor(limiter), MatchFunc(healthSkip)),
		),
		grpc.ChainStreamInterceptor(
			StreamServerInterceptor(ratelimit.StreamServerInterceptor(limiter), MatchFunc(healthSkip)),
		),
	)
}
