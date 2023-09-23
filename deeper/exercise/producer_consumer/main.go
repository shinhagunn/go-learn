package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numProducers  = 3
	numConsumers  = 2
	maxBufferSize = 5
)

func main() {
	rand.Seed(time.Now().UnixNano())

	buffer := make(chan int, maxBufferSize)
	done := make(chan bool)
	var wg sync.WaitGroup

	// Goroutines for producers
	for i := 0; i < numProducers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				value := rand.Intn(100)
				fmt.Printf("Producer %d produced %d\n", id, value)
				buffer <- value // Send value to the buffer channel
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
			}
		}(i)
	}

	// Goroutines for consumers
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				value := <-buffer // Receive value from the buffer channel
				fmt.Printf("Consumer %d consumed %d\n", id, value)
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
			}
		}(i)
	}

	// Signal when done
	go func() {
		time.Sleep(time.Second * 5) // Run the program for 5 seconds
		done <- true
	}()

	wg.Wait()     // Wait for all producers and consumers to finish
	close(buffer) // Close the buffer channel
	<-done        // Wait for the signal that the program is done
}
