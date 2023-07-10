package recovery

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"runtime/debug"
	"time"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		st := time.Now()
		defer func() {
			if recoverError := recover(); recoverError != nil {
				log.Error().Str("method", info.FullMethod).Interface("recovery", recoverError).Bytes("stack", debug.Stack()).Interface("req", req).Interface("resp", resp).Err(err).TimeDiff("runtime", time.Now(), st).Send()
			} else {
				log.Info().Str("method", info.FullMethod).Interface("req", req).Interface("resp", resp).Err(err).TimeDiff("runtime", time.Now(), st).Send()
			}
		}()
		return handler(ctx, req)
	}
}

func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		st := time.Now()
		defer func() {
			if recoverError := recover(); recoverError != nil {
				log.Error().Str("method", info.FullMethod).Interface("recovery", recoverError).Bytes("stack", debug.Stack()).Err(err).TimeDiff("runtime", time.Now(), st).Send()
			} else {
				log.Info().Str("method", info.FullMethod).Err(err).TimeDiff("runtime", time.Now(), st).Send()
			}
		}()
		return handler(srv, stream)
	}
}
