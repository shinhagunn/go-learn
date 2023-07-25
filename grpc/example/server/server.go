package main

import (
	"context"
	"fmt"
	"net"

	"github.com/shinhagunn/go-learn/grpc/example/calculator"
	"google.golang.org/grpc"
)

type calculatorServer struct {
	calculator.UnimplementedCalculatorServer
}

func (s *calculatorServer) Add(ctx context.Context, req *calculator.AddRequest) (*calculator.AddResponse, error) {
	result := req.GetNum1() + req.GetNum2()
	return &calculator.AddResponse{
		Result: result,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	calculator.RegisterCalculatorServer(server, &calculatorServer{})

	fmt.Println("Server started on port 50051")
	if err = server.Serve(listener); err != nil {
		panic(err)
	}
}
