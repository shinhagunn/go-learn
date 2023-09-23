package main

import (
	"context"
	"fmt"

	hello "github.com/shinhagunn/go-learn/grpc/gateway/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)

	res, err := client.Hello(context.TODO(), &hello.HelloRequest{Name: "hoang"})
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
