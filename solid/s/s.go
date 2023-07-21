package s

import "fmt"

// SOLID
// S - Single Responsibility Principle
// A class should have one and only one reason to change, meaning that a class should have only one job

// Bad
type StudentBad struct {
	Name string
	Age  int
}

func (s StudentBad) PrintInfo() string {
	return fmt.Sprintf("Name: %s - Age: %d", s.Name, s.Age)
}

func (s StudentBad) SaveDatabase() bool {
	return true
}

// Good
type StudentGood struct {
	Name string
	Age  int
}

func (s StudentGood) PrintInfo() string {
	return fmt.Sprintf("Name: %s - Age: %d", s.Name, s.Age)
}

type Database struct {
	// data
}

func (d Database) Save() bool {
	return true
}
