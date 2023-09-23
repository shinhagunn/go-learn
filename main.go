package main

import "fmt"

func a() {
	i := 0
	
	defer func() {
		fmt.Println(i)
	}()
	i++

	return
}

func main() {
	a()
}
