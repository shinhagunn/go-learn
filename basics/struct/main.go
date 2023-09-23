package main

import "fmt"

// Example
type student struct {
	name string
	age  int
}

func (s student) GetInfo() {
	fmt.Printf("%s: %d\n", s.name, s.age)
}

func (s *student) SetName(name string) {
	s.name = name
}

func main() {
	fmt.Println("Hello Struct!")

	a := student{
		name: "Hieu",
		age:  18,
	}

	b := a
	b.age = 20

	a.GetInfo()
	b.GetInfo()
}
