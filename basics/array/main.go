package main

import "fmt"

func RemoveElementFromIndex(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func main() {
	fmt.Println("Hello Array!")

	a := []int{0, 1, 2, 3, 4, 5}
	b := RemoveElementFromIndex(a, 3)
	fmt.Println(b)
}
