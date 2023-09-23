package main

import (
	"fmt"
	"time"
)

// Declare channel
// c := make(chan int) or c := make(chan int, 0) -> this is unbuffer channel integers have capacity equal 0
// c := make(chan int, 1) -> this is buffer channel integers have capacity equal 1

// Here is an example of channel

const students = 5
const numberOfTasks = 1000

func giveTasks(c chan int) {
	// Remember close channel when done tasks
	defer func() {
		close(c)
		fmt.Println("Out of task")
	}()

	for i := 1; i <= numberOfTasks; i++ {
		c <- i
		fmt.Printf("Give task number: %d\n", i)
	}
}

func reciveTasks(c chan int, name int) {
	// Can use for range to check a channel is closed
	for v := range c {
		fmt.Printf("Student number %d recive task: %d\n", name, v)
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Printf("Student number %d done\n", name)
}

func main() {
	c := make(chan int, 100)
	go giveTasks(c)

	for i := 1; i <= students; i++ {
		go reciveTasks(c, i)
	}

	time.Sleep(5 * time.Minute)
}
