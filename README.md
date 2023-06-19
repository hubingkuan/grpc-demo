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

## 注意事项

**proto文件中的package和option go_package的区别:**

* package: 属于proto文件自身的范围定义的包名 其他proto引用该proto时使用该包名import
* option go_package: 生成的go文件的包名,路径,包名和路径之间用分号隔开
* --go_out=paths=source_relative:. 参数是为了让加了 option go_package 声明的 proto 文件可以将 go 代码编译到与其同目录

**不同包之间的 proto 文件不可以循环依赖，这会导致生成的 go 包之间也存在循环依赖，导致 go 代码编译不通过**

**同属于一个包内的 proto 文件之间的引用也需要声明 import**

grpc中间件参考: [中间件](https://github.com/grpc-ecosystem/go-grpc-middleware)