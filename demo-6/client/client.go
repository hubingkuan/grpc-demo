package main

import (
	"context"
	"fmt"

	"grpc-demo/demo-6/example"

	"github.com/apache/thrift/lib/go/thrift"
)

func main() {

	conf := &thrift.TConfiguration{}
	transport := thrift.NewTSocketConf("localhost:8080", conf)
	defer transport.Close()
	// 创建客户端协议
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(conf)
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
