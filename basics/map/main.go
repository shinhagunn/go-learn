package main

import "fmt"

func main() {
	fmt.Println("Hello Map!")

	m := make(map[string]bool)
	m["a"] = true
	m["b"] = false

	// delete(m, "b")
	v, ok := m["a"]
	if ok {
		fmt.Println(v)
	}

	fmt.Println(m)
}
