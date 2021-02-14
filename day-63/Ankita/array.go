package main

import "fmt"

func mapArray(arr []int, callback func(int) int) []int {
	newArray := make([]int, len(arr))
	for index, value := range arr {
		newArray[index] = callback(value)
	}

	return newArray
}

func main() {
	square := func(x int) int { return x * x }
	fmt.Println(mapArray([]int{1, 2, 3, 4, 5}, square)) // prints [1 4 9 16 25]
}
