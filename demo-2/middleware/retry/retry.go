package retry

import (
	"context"
	"google.golang.org/grpc"
)

func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(parentCtx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// 简易版本 后续可添加参数重试次数  可忽略的错误 重试间隔时间  重试回调等等
		var lastErr error
		for attempt := uint(0); attempt < 5; attempt++ {
			lastErr = invoker(parentCtx, method, req, reply, cc, opts...)
			if lastErr == nil {
				return nil
			}
		}
		return lastErr
	}
}
