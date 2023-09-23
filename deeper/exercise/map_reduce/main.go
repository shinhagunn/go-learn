package main

import (
	"fmt"
	"sync"
)

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	splits := 4
	c := make(chan int)

	// map
	var wg sync.WaitGroup
	for i := 0; i < splits; i++ {
		start := i * len(data) / splits
		end := (i + 1) * len(data) / splits

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()

			sum := 0
			for i := start; i < end; i++ {
				sum += data[i]
			}

			c <- sum
		}(start, end)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	result := 0
	for s := range c {
		result += s
	}

	fmt.Printf("Total sum: %d\n", result)
}
