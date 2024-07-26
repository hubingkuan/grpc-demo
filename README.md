# 项目目录

1. demo-1: 最简单的grpc服务 server中的proto引用client中的proto的message
2. demo-2: 拦截器(token校验、日志记录、校验参数、限流、重试、selector、recovery、timeout、keepalive、监控(
   prometheus、opentracing))
   、metadata 客户端 服务端互传数据、proto生成脚本(validate、doc、inject tag、grpc)
3. demo-3: 流式 grpc示例 流式拦截器
4. demo-4: 服务注册与服务发现  (etcd官方包实现resolver)
5. demo-5: 服务注册与服务发现  (etcd、zk、nacos 实现自定义resolver 服务注册+服务发现+服务配置)
6. demo-6: 简单的thrift服务
7. demo-7: opentelemetry+jaeger简单实现(日志(ELK):应用程序内打的日志+指标(prometheus):
   提供系统运行状况视图,是否在期望的边界内运行+跟踪(jaeger):可视化请求在整个系统中移动时的进度)
8. demo-8: grpc-gateway+openapiv2(将grpc服务转换成http服务)

metadata参考:[metadata](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-metadata.md#unary-call)

grpc中间件参考: [中间件](https://github.com/grpc-ecosystem/go-grpc-middleware)

grpc网关参考: [grpc-gateway](https://grpc-ecosystem.github.io/grpc-gateway/)

thrift参考: [thrift](https://thrift.apache.org/)

## Grpc基础

1、下载protobuf:https://github.com/protocolbuffers/protobuf/releases/tag/v23.1 (注意版本和公司保持一致 低版本大概是3.17.3
高版本与低版本生成代码的命令不一致 protoc-gen-go的版本也需要变化 go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1)

* 设置系统环境变量: D:\protobuf\bin
* 查看proto版本: protoc--version

2、安装核心库:

* go get google.golang.org/grpc
* go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
* go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
* go install github.com/bufbuild/protoc-gen-validate@latest ([pb验证参数](https://github.com/bufbuild/protoc-gen-validate/blob/main/docs.md))
* go install github.com/favadi/protoc-go-inject-tag@latest (pb生成struct tag 参考login.proto)
* go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest (pb生成接口文档)
* go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway (grpc反向代理生成http服务)
* go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 (pb生成swagger文档)

3、编写proto文件 规范

* 包名package 小写
* message和service使用驼峰命名 rpc方法名使用驼峰命名 message里面的字段命名采用小写字母+下划线分割

```go
message SongServerRequest {
optional string song_name = 1;
}
```

* enum使用驼峰命名 字段命名采用大写字母+下划线分割 使用分号结尾 默认值使用第一个枚举值

```go
enum Foo {
FOO_UNSPECIFIED = 0;
FOO_FIRST_VALUE = 1;
FOO_SECOND_VALUE = 2;
}
```

* 重复字段使用复数名称 repeated string keys =1;

4、protobuf的数据类型

```go
protoType       GoType
double          float64
float           float32
int32           int32
int64           int64
uint32          uint32
uint64          uint64
sint32          int32
sint64          int64
fixed32         uint32
fixed64         uint64
sfixed32        int32
sfixed64        int64
bool            bool
string          string
bytes           []byte
enum
message
map<T, K>
repeated T
oneof:  如果消息中有很多可选字段 并且同时最多只能有一个字段被设置为非默认值 那么可以使用oneof (多选一 oneof内的字段不能用Repeated修饰)
```

5、proto命令使用

#### 新版生成代码命令

1. protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative
   friend/friend.proto

* .代表生成的目录

2. protoc --validate_out="lang=go,paths=source_relative:./"  friend/friend.proto

3. protoc --doc_out=./doc --doc_opt=html,friend.html friend/friend.proto

4. protoc-go-inject-tag -input=./*/*.pb.go

* [pb文件结构体tag注入](https://github.com/favadi/protoc-go-inject-tag)


5. protoc --grpc-gateway_out=./ --grpc-gateway_opt paths=source_relative --grpc-gateway_opt logtostderr=true
   --grpc-gateway_opt generate_unbound_methods=true friend/friend.proto

* --grpc-gateway_opt standalone=true 与源proto文件分开生成 生成一个单独的网关包 一般不设置
* --grpc-gateway_opt logtostderr=true 代表将日志输出到标准错误输出
* --grpc-gateway_opt generate_unbound_methods=true 代表对于没有定义options的rpc方法也自动生成映射http方法,默认生成的方法是post,url路径是
  /全路径service name/method name 也就是 /package name.service name/method name
* 需要将google/api/annotations.proto和google/api/http.proto 放在protoc同级目录下

6. protoc --openapiv2_out=./ --openapiv2_opt=logtostderr=true --openapiv2_opt=generate_unbound_methods=true
   --openapiv2_opt=preserve_rpc_order=true --openapiv2_opt=allow_merge=true,merge_file_name=demo   */*.proto

* --openapiv2_opt=generate_unbound_methods 将未生成option的rpc方法也生成OpenAPI定义
* --openapiv2_opt=preserve_rpc_order=true 保留rpc方法的顺序
* --openapiv2_opt=allow_merge=true,merge_file_name=foo 将不同的输入合并到单个 OpenAPI 文件中 文件名是foo.swagger.json
* 需要将protoc-gen-openapiv2/options/annotations.proto放在protoc同级目录下
* // 生成*.swagger.json文件的一些默认设置  
  option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = {  
  host: "localhost:8080"
  base_path: ""
  info: {  
  title: "friend api docs";
  version: "v1.0";  
  };  
  // 默认为HTTPS，根据实际需要设置  
  schemes: HTTP;  
  // 显示扩展文档  
  external_docs: {  
  url: "https://baidu.com";  
  description: "描述信息";  
  }
  };

  rpc Login(LoginRequest) returns (LoginReply) {
  option (google.api.http) = {
  post: "/api/v1/login"
  body: "*"
  };
  option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
  summary: "登录",
  description: "登录",
  tags: "user",
  parameters: {
  headers: {
  name: "X-Foo";
  description: "Foo Header";
  type: STRING,
  required: true;
  };
  headers: {
  name: "X-Bar";
  description: "Bar Header";
  type: NUMBER,
  };
  };
  };
  }
  }

#### 旧版生成代码命令

protoc --go_out=plugins=grpc:./ --go_opt=paths=source_relative friend/friend.proto

**_注意事项_**

1. proto文件中的package和option go_package的区别:

  * package: 属于proto文件自身的范围定义的包名 其他proto引用该proto时使用该包名import
  * option go_package: 生成的go文件的包名,路径,包名和路径之间用分号隔开
  * --go_out=paths=source_relative:. 参数是为了让加了 option go_package 声明的 proto 文件可以将 go 代码编译到与其同目录

2. 不同包之间的 proto 文件不可以循环依赖,这会导致生成的 go 包之间也存在循环依赖,导致 go 代码编译不通过
3. 同属于一个包内的 proto 文件之间的引用也需要声明 import

## 传参

1. 如何构建metadata

```go 
md := metadata.Pairs(
"key1", "val1",
"key1", "val1-2",
"key2", "val2",
)

md := metadata.New(map[string]string{"key1": "val1", "key2": "val2"})

// 注意:所有键将自动转换为小写, 因此,“key1”和“kEy1”将是相同的键,它们的值将合并到同一个列表中
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

## Thrift基础

1、下载Thrift Golang库: go get github.com/apache/thrift/lib/go/thrift

2、下载thrift编译器:https://dlcdn.apache.org/thrift/0.18.1/

* 将下载的exe改名为thrift.exe
* 设置系统环境变量:D:\thrift
* 查看thrift版本 thrift -version

3、编写thrift文件 规范

4、thrift的数据类型

```go
// 基本类型
bool:布尔型, 4位
byte:带符号整数,8位
i16:带符号整数, 18位
i32:带符号整数, 32位
i64:带符号整数, 64位
double:64位浮点型
string:UTF-8编码的字符串
// 特殊类型
binary:未经编码的字节流
// 结构体
struct:公共对象, 不能继承
struct test{
1: string name
}
// 枚举
enum test{
OK = 0,
Fail = 1
}
// 容器
list<T>:    有序列表
set<T>:        无序集合
map<T, K>:    映射数据
// 异常类型
exception:
// 服务类型
service:对应服务的类
```

5、生成go文件命令:

* thrift -out .. --gen go example.thrift : 在同级目录下生成golang的包 生成的format_data-remote是生成的测试文件