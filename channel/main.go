package main

import (
	"fmt"
	"time"
)

// Basic
// Declare channel
// c := make(chan int) or c := make(chan int, 0) -> this is unbuffer channel integers have capacity equal 0
// c := make(chan int, 1) -> this is buffer channel integers have capacity equal 1

// Here is an example of channel

const students = 5
const numberOfTasks = 1000

func giveTasks(c chan int) {
	for i := 1; i <= numberOfTasks; i++ {
		c <- i
		fmt.Printf("Give task number: %d\n", i)
		time.Sleep(10 * time.Millisecond)
	}
}

func reciveTasks(c chan int, name int) {
	for i := 1; i <= numberOfTasks; i++ {
		fmt.Printf("Student number %d recive task: %d\n", name, <-c)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	c := make(chan int, 100)
	go giveTasks(c)

	for i := 1; i <= students; i++ {
		go reciveTasks(c, i)
	}

	time.Sleep(5 * time.Minute)
}
