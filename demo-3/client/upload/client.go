package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-demo/demo-3/proto/upload"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := upload.NewUploadServiceClient(conn)
	stream, err := client.UploadImage(context.Background())

	imagePath := fmt.Sprintf("%s/laptop.jpg", "../tmp")
	file, err := os.Open(imagePath)
	defer file.Close()

	imageType := filepath.Ext(imagePath)
	req := &upload.UploadImageRequest{
		Data: &upload.UploadImageRequest_ImageInfo{
			ImageInfo: &upload.ImageInfo{
				LaptopId:  "1",
				ImageType: imageType,
			},
		},
	}

	err = stream.Send(req)

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	size := 0

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}

		size += n

		req := &upload.UploadImageRequest{
			Data: &upload.UploadImageRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
	}

	res, err := stream.CloseAndRecv()

	savedImagePath := fmt.Sprintf("%s/%s%s", "../tmp", res.GetId(), imageType)
	fmt.Println(savedImagePath)
}
