package main

import "fmt"

// In Go, function like other languages

// Return values
func CalValues(a int, b int) (int, int) {
	return a + b, a - b
}

// Function named
func CalNamed(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	fmt.Println("Hello Function")
}
