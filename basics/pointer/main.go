package main

import "fmt"

func squareArray(arr []int) (result []int) {
	result = []int{}

	for i := 0; i < len(arr); i++ {
		result = append(result, arr[i]*arr[i])
	}

	return
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(squareArray(arr))
}
