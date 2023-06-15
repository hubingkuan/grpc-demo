1、下载protobuf:https://github.com/protocolbuffers/protobuf/releases/tag/v23.1
设置系统环境变量:D:\protobuf\bin

2、安装核心库:
go get google.golang.org/grpc 
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

3、编写proto文件

4、生成go文件命令:
   protoc --go_out=. hello.proto  生成go语言文件
   protoc --go-grpc_out=. hello.proto  生成rpc文件
 .代表生成的目录  hello.proto代表根据这个proto生成文件

grpc中间件参考: [中间件](https://github.com/grpc-ecosystem/go-grpc-middleware)


其他grpc插件:

pb验证参数插件:https://github.com/bufbuild/protoc-gen-validate


pb生成接口文档插件:https://github.com/pseudomuto/protoc-gen-doc

protoc --validate_out="lang=go:./gen"  --doc_out=./doc --doc_opt=html,index.html --go_out=./gen/ proto/*.proto