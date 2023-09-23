package main

import (
	"context"
	"fmt"
	"net"

	hello "github.com/shinhagunn/go-learn/grpc/gateway/proto"
	"google.golang.org/grpc"
)

type helloServer struct {
	hello.UnimplementedHelloServiceServer
}

func (h *helloServer) Hello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	fmt.Println("Receive: ", in.GetName())
	return &hello.HelloResponse{Msg: "Name: " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, &helloServer{})

	fmt.Println("Start server at: ", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
