# 基础
1、下载protobuf:https://github.com/protocolbuffers/protobuf/releases/tag/v23.1
设置系统环境变量:D:\protobuf\bin

2、安装核心库:

* go get google.golang.org/grpc
* go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
* go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
* go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest (pb生成接口文档)
* go install github.com/bufbuild/protoc-gen-validate@latest (pb验证参数)

3、编写proto文件

4、生成go文件命令:

* protoc --go_out=. hello.proto 生成go语言文件
* protoc --go-grpc_out=. hello.proto 生成rpc文件
  .代表生成的目录 hello.proto代表根据这个proto生成文件
* protoc --validate_out="lang=go:./gen"  --doc_out=./doc --doc_opt=html,index.html --go_out=./gen/ proto/*.proto (
  生成pb的接口文档html 以及go文件生成)
* protoc --go_out=plugins=grpc,paths=source_relative:. hello.proto (将rpc和go文件合成在同一个文件下
  并在相对目录下生成文件)
* protoc -I ../../../ -I ./ --go_out=plugins=grpc,paths=source_relative:. hello2.proto  (-I 指定搜索proto文件的目录
  ../../../就到了gopath下 另外hello2.proto的import是从项目根路径开始 go mod定义  )

**_注意事项_**

1. proto文件中的package和option go_package的区别:

    * package: 属于proto文件自身的范围定义的包名 其他proto引用该proto时使用该包名import
    * option go_package: 生成的go文件的包名,路径,包名和路径之间用分号隔开
    * --go_out=paths=source_relative:. 参数是为了让加了 option go_package 声明的 proto 文件可以将 go 代码编译到与其同目录

2. 不同包之间的 proto 文件不可以循环依赖，这会导致生成的 go 包之间也存在循环依赖，导致 go 代码编译不通过
3. 同属于一个包内的 proto 文件之间的引用也需要声明 import

# 传参

1. 如何构建metadata

```go 
md := metadata.Pairs(
"key1", "val1",
"key1", "val1-2",
"key2", "val2",
)

md := metadata.New(map[string]string{"key1": "val1", "key2": "val2"})

// 注意：所有键将自动转换为小写， 因此，“key1”和“kEy1”将是相同的键，它们的值将合并到同一个列表中
``` 

2. 客户端发送metadata

```go
// 建立连接
client := pb.NewSayHelloClient(conn)
// 这里创建metadata给服务端(所有的key都会转换成小写)
md := metadata.Pairs("appid", "maguahu", "appkey", "12345")
ctx := metadata.NewOutgoingContext(context.Background(), md)
// 也可以使用ctx := metadata.AppendToOutgoingContext(ctx, "k1", "v1", "k1", "v2", "k2", "v3")
// 执行rpc调用 传入ctx(包含metadata)
response, err := client.SayHello(ctx, &pb.HelloRequest{RequestName: "maguahu"})
// 注意: 一元调用和流式调用 客户端都是使用ctx发送metadata
```

3. 客户端接受metadata

```go
// 一元调用
var header, trailer metadata.MD
r, err := client.SomeRPC(
ctx,
someRequest,
grpc.Header(&header), // 接收的header放在这里   本质是opts ...grpc.CallOption  
grpc.Trailer(&trailer), // 接收的trailer放这里
)
fmt.Println(header.Get("key")) // 打印从服务端这边得到的md中定义的key

// 流式调用
stream, err := client.SomeStreamingRPC(ctx)
// retrieve header
header, err := stream.Header()
// retrieve trailer
trailer := stream.Trailer()
```

4. 服务端发送metadata

```go
// 一元调用:
func (s *server) SomeRPC(ctx context.Context, in *pb.someRequest) (*pb.someResponse, error) {
// 创建并设置 header
header := metadata.Pairs("header-key", "val")
grpc.SendHeader(ctx, header)
// 创建并设置 trailer
trailer := metadata.Pairs("trailer-key", "val")
grpc.SetTrailer(ctx, trailer)
}

// 流式调用:
func (s *server) SomeStreamingRPC(stream pb.Service_SomeStreamingRPCServer) error {
// create and send header
header := metadata.Pairs("header-key", "val")
stream.SendHeader(header)
// create and set trailer
trailer := metadata.Pairs("trailer-key", "val")
stream.SetTrailer(trailer)
}

```

5. 服务端接受metadata

```go
// 一元调用
md, ok := metadata.FromIncomingContext(ctx)

// 流式
md, ok := metadata.FromIncomingContext(stream.Context())
```

6. 服务端grpc包装自身数据 创建新的上下文(拦截器常用)

```go
metadata.NewIncomingContext(ctx, md)
```

7. 自身grpc获取metadata

```go
// 获取已有的metadata与新的metadata合并
send, _ := metadata.FromOutgoingContext(ctx)
newMD := metadata.Pairs("k3", "v3")
ctx = metadata.NewOutgoingContext(ctx, metadata.Join(send, newMD))
```

# 项目目录

1. demo-1: 最简单的grpc服务 server中的proto引用client中的proto的message
2. demo-2: 拦截器(token校验、日志记录、校验参数、限流、重试、recovery、timeout、keepalive、监控(prometheus、opentracing))
   、metadata 客户端 服务端互传数据、proto生成脚本(validate、doc、grpc)
3. demo-3: 流式 grpc示例 流式拦截器
4. demo-4: 服务注册与服务发现  (etcd简易版本)
5. demo-5: 服务注册与服务发现  (etcd完整版本)

详情可参考:[metadata](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-metadata.md#unary-call)

grpc中间件参考: [中间件](https://github.com/grpc-ecosystem/go-grpc-middleware)