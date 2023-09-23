package main

import (
	"fmt"
	"sync"
	"time"
)

const numbers = 300
const worker = 5

func giveNumber(c chan<- int) {
	defer func() {
		close(c)
		fmt.Println("Done give all number")
	}()

	for i := 0; i < numbers; i++ {
		c <- i
		fmt.Printf("Give number: %d\n", i)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	c := make(chan int, 100)
	go giveNumber(c)

	var wg sync.WaitGroup
	wg.Add(worker)
	for i := 0; i < worker; i++ {
		go func(number int) {
			name := fmt.Sprintf("worker %d", number)

			for v := range c {
				fmt.Printf("%s take number: %d\n", name, v)
				time.Sleep(100 * time.Millisecond)
			}

			fmt.Printf("%s take number done!\n", name)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
