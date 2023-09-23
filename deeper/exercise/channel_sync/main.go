package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 1; i <= 5; i++ {
			c <- i
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 6; i <= 10; i++ {
			c <- i
		}
	}()

	go func() {
		for number := range c {
			fmt.Println(number)
		}
	}()

	wg.Wait()
	close(c)
}
