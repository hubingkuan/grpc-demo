package interceptor

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func rpcString(v interface{}) string {
	if s, ok := v.(interface{ String() string }); ok {
		return s.String()
	}
	return fmt.Sprintf("%+v", v)
}

func RpcClientInterceptor(
	ctx context.Context,
	method string,
	req, resp interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) (err error) {
	if ctx == nil {
		return errors.New("call rpc request context is nil")
	}
	// 包装ctx传值 metadata.NewOutgoingContext
	ctx, err = getRpcContext(ctx, method)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	fmt.Println(ctx, "get rpc ctx success", "conn target", cc.Target())
	err = invoker(ctx, method, req, resp, cc, opts...)
	if err == nil {
		fmt.Println(ctx, "rpc client resp", "funcName", method, "resp", rpcString(resp))
		return nil
	}
	fmt.Println(ctx, "rpc resp error", err)
	return err
}

func getRpcContext(ctx context.Context, method string) (context.Context, error) {
	md := metadata.Pairs()
	/*	// 必要参数传值
		operationID, ok := ctx.Value("OperationID").(string)
		if !ok {
			fmt.Println(ctx, "ctx missing operationID", errors.New("ctx missing operationID"), "funcName", method)
			return nil, errors.New("ctx missing operationID")
		}
		md.Set("OperationID", operationID)
		// 非必要参数传值
		connID, ok := ctx.Value("ConnID").(string)
		if ok {
			md.Set("ConnID", connID)
		}*/
	return metadata.NewOutgoingContext(ctx, md), nil
}
