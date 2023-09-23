package main

import "fmt"

// Example Generic Type
type Response[T any] struct {
	Data []T
}

// Example Generic Interface
type IArea[T any] interface {
	Area() T
}

// Example Generic Func
func Sum[T string | int | int64 | float64](a T, b T) T {
	return a + b
}

func main() {
	fmt.Println("Hello Generics!")

	fmt.Println(Sum("a", "b"))
	fmt.Println(Sum(3, 4))
	fmt.Println(Sum(5.5, 6.5))
}
