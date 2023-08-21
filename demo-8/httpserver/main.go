package main

import (
	"context"
	"flag"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"grpc-demo/demo-8/proto/friend"
	"net/http"
	"strconv"
)

var (
	grpcServerEndpoint = flag.String("endpoint", "localhost:8082", "grpc server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	mux := runtime.NewServeMux(
		// 将请求方法和路径添加到grpc元数据中
		runtime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
			md := make(map[string]string)
			if method, ok := runtime.RPCMethod(ctx); ok {
				md["method"] = method
			}
			if pattern, ok := runtime.HTTPPathPattern(ctx); ok {
				md["pattern"] = pattern
			}
			return metadata.New(md)
		}),
		// 将特殊的http的请求头 并保留默认的映射规则 加入到grpc的metadata中
		runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
			switch key {
			case "X-Request-Id":
				return key, true
			default:
				return runtime.DefaultHeaderMatcher(key)
			}
		}),
		// 改变响应消息 或设置响应头  相当于后置响应处理器
		runtime.WithForwardResponseOption(func(ctx context.Context, w http.ResponseWriter, responseMessage proto.Message) error {
			// 可以从响应中拿到player_id 放到响应头中
			t, ok := responseMessage.(*friend.RadarSearchPlayerInfo)
			if ok {
				w.Header().Set("player_id", strconv.Itoa(int(t.GetPlayerId())))
			}
			// 也可以从服务端设置的元数据中拿到metadata 数据
			md, ok := runtime.ServerMetadataFromContext(ctx)
			if !ok {
				return nil
			} else {
				if code := md.HeaderMD.Get("code"); len(code) > 0 {
					code, _ := strconv.Atoi(code[0])
					// 修改http状态码
					w.WriteHeader(code)
				}
			}
			return nil
		}),
		//  自定义路由错误
		// runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
		//
		// }),
		// 自定义序列化
		// runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		// 	MarshalOptions: protojson.MarshalOptions{
		// 		Indent:    "  ",
		// 		Multiline: true, // Optional, implied by presence of "Indent".
		// 	},
		// 	UnmarshalOptions: protojson.UnmarshalOptions{
		// 		DiscardUnknown: true,
		// 	},
		// }),
	)
	options := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := friend.RegisterFriendHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, options)
	if err != nil {
		return err
	}
	// 可以添加自定义路由处理器
	mux.HandlePath("GET", "/v1/hello/{name}", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Write([]byte("hello " + pathParams["name"]))
	})
	glog.Info("Starting http server...")
	return http.ListenAndServe(":8081", mux)
}

// 启动http server服务 充当rpc代理客户端
func main() {
	flag.Parse()
	// 通过命令行传递参数 –-log_dir指定日志文件的存放目录，默认为os.TempDir()
	defer glog.Flush()
	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
