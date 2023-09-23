package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		time.Sleep(2 * time.Minute)
		c <- 1
	}()

	select {
	case value := <-c:
		fmt.Println(value)
	case <-time.After(2 * time.Second):
		fmt.Println("Qua thoi gian")
	}
}
