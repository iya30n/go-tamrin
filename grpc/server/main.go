package main

import (
	"context"
	"grpc/api/helloworld"
	"log"
	"net"

	"google.golang.org/grpc"
)

type HelloworldServer struct {
	helloworld.UnimplementedGreeterServer
}

func (hs HelloworldServer) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	return &helloworld.HelloResponse{Message: "hello " + in.GetName()}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalln("Listen failed: ", err)
	}

	server := grpc.NewServer()

	helloworld.RegisterGreeterServer(server, &HelloworldServer{})

	log.Printf("Server is listening on port: %v", listener.Addr())

	if err := server.Serve(listener); err != nil {
		log.Fatalln("Failed to server: ", err)
	}
}