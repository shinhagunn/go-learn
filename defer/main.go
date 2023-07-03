package main

import (
	"fmt"
)

// A function will run statements in defer when return
// Defer have 3 rules:
// 1. A deferred function’s arguments are evaluated when the defer statement is evaluated
// 2. Deferred function calls are executed in LIFO
// 3. Deferred functions may read and assign to the returning function’s named return values.

// Result equal 2
func a() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	fmt.Println("Hello defer!")
}
