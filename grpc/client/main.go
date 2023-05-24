package main

import (
	"context"
	"grpc/api/helloworld"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	grpcClient, err := grpc.Dial("127.0.0.1:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err !=nil {
		log.Fatalf("Connection failed: %v", err)
	}
	defer grpcClient.Close()

	helloClient := helloworld.NewGreeterClient(grpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := helloClient.SayHello(ctx, &helloworld.HelloRequest{Name: "sina"})
	if err != nil {
		log.Fatalf("Request Failed: %v", err)
	}

	log.Println(res.GetMessage())
}