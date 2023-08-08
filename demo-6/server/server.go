package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"grpc-demo/demo-6/example"
	"log"
)

type FormatDataImpl struct{}

func (f FormatDataImpl) SayHello(ctx context.Context, name string) (_r *example.Person, _err error) {
	age := int32(25)
	return &example.Person{
		Name: name,
		Age:  &age,
	}, nil
}

const (
	HOST = "localhost"
	PORT = "8080"
)

func main() {

	// 创建基于socket的传输通道
	serverTransport, err := thrift.NewTServerSocket(HOST + ":" + PORT)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	// 创建thrift传输器和协议
	transportFactory := thrift.NewTTransportFactory()
	confN := &thrift.TConfiguration{}
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(confN)

	// 创建thrift服务器
	handler := &FormatDataImpl{}
	processor := example.NewFormatDataProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Running at:", HOST+":"+PORT)
	err = server.Serve()
	if err != nil {
		log.Fatalln("Error:", err)
	}
}
