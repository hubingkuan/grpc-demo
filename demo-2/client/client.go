package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-demo/demo-2/proto/hello"
	"log"
)

func main() {
	conn, _ := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	client := hello.NewGreeterClient(conn)
	res, err := client.SayHello(context.Background(), &hello.Person{
		// 这里故意传递一个不符合规范的值
		Id:    988,
		Email: "fanqiechaodan@fanqiechaodan.com",
		Name:  "番茄炒蛋",
		Home: &hello.Person_Location{
			Lat: 23,
			Lng: 45,
		},
	})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(res.Id)
}
