package timeout

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"time"
)

func UnaryClientInterceptor(timeout time.Duration) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		st := time.Now()
		timedCtx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		err := invoker(timedCtx, method, req, reply, cc, opts...)
		if err != nil {
			log.Error().Str("method", method).Interface("req", req).Interface("resp", reply).Err(err).TimeDiff("runtime", time.Now(), st).Send()
		} else {
			log.Info().Str("method", method).Interface("req", req).Interface("resp", reply).Err(err).TimeDiff("runtime", time.Now(), st).Send()
		}
		return err
	}
}
