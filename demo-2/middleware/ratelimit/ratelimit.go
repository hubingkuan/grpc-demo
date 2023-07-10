package ratelimit

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Limiter interface {
	Limit(ctx context.Context) error
}

// 或者使用开源框架github.com/juju/ratelimit实现

// 以channel为核心的限流器
type ChannelTokenLimiter struct {
	rate         int
	tokenChannel chan struct{}
}

func NewChannelTokenLimiter(rate int) *ChannelTokenLimiter {
	limiter := &ChannelTokenLimiter{
		rate:         rate,
		tokenChannel: make(chan struct{}, rate),
	}

	go func() {
		for range time.Tick(1 * time.Second) {
			for i := 0; i < rate; i++ {
				limiter.tokenChannel <- struct{}{}
			}
		}
	}()
	return limiter
}

func (channelTokenLimiter *ChannelTokenLimiter) Limit(ctx context.Context) error {
	select {
	case <-channelTokenLimiter.tokenChannel:
		return nil
	default:
		return fmt.Errorf("reached Rate-Limiting %d", channelTokenLimiter.rate)
	}
	return nil
}

func UnaryServerInterceptor(limiter Limiter) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if err := limiter.Limit(ctx); err != nil {
			return nil, status.Errorf(codes.ResourceExhausted, "%s is rejected by grpc_ratelimit middleware, please retry later. %s", info.FullMethod, err)
		}
		return handler(ctx, req)
	}
}

func StreamServerInterceptor(limiter Limiter) grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if err := limiter.Limit(stream.Context()); err != nil {
			return status.Errorf(codes.ResourceExhausted, "%s is rejected by grpc_ratelimit middleware, please retry later. %s", info.FullMethod, err)
		}
		return handler(srv, stream)
	}
}

func UnaryClientInterceptor(limiter Limiter) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if err := limiter.Limit(ctx); err != nil {
			return status.Errorf(codes.ResourceExhausted, "%s is rejected by grpc_ratelimit middleware, please retry later. %s", method, err)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func StreamClientInterceptor(limiter Limiter) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		if err := limiter.Limit(ctx); err != nil {
			return nil, status.Errorf(codes.ResourceExhausted, "%s is rejected by grpc_ratelimit middleware, please retry later. %s", method, err)
		}
		return streamer(ctx, desc, cc, method, opts...)
	}
}
