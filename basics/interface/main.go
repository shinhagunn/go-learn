package main

import "fmt"

type IArea interface {
	Area() float64
}

type Square struct {
	Width float64
}

func (s Square) Area() float64 {
	return s.Width * s.Width
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (s Rectangle) Area() float64 {
	return s.Width * s.Height
}

func main() {
	fmt.Println("Hello Interface!")

	var x IArea = Square{5.5}
	var y IArea = Rectangle{5, 6}

	fmt.Println(x.Area())
	fmt.Println(y.Area())
}
