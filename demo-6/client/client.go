package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"grpc-demo/demo-6/example"
)

func main() {

	conf := &thrift.TConfiguration{}
	transport := thrift.NewTSocketConf("localhost:8080", conf)

	// 创建客户端协议
	confN := &thrift.TConfiguration{}
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(confN)
	client := example.NewFormatDataClientFactory(transport, protocolFactory)

	// 打开连接
	if err := transport.Open(); err != nil {
		panic(err)
	}
	defer transport.Close()

	// 调用远程服务
	if person, err := client.SayHello(context.Background(), "world"); err != nil {
		panic(err)
	} else {
		fmt.Println(person)
	}
}
