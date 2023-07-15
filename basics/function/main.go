package function

import "fmt"

// In Go, function like other languages

// Return values
func CalValues(a int, b int) (int, int) {
	return a + b, a - b
}

// Function named
func CalNamed(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func main() {
	fmt.Println("Hello Function")
}
