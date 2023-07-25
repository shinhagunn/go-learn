package main

import (
	"context"
	"fmt"

	"github.com/shinhagunn/go-learn/grpc/example/calculator"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := calculator.NewCalculatorClient(conn)

	num1 := int32(10)
	num2 := int32(20)

	res, err := client.Add(context.Background(), &calculator.AddRequest{
		Num1: num1,
		Num2: num2,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
