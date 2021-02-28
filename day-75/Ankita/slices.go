package main

import "fmt"

func main() {
	var i [5]int = [5]int{1, 2, 3, 4, 5}
	var slicedI []int = i[1:3]
	fmt.Println("slicedI = ", slicedI)
	fmt.Println("i = ", i)

	slicedI = []int{7, 8}
	fmt.Println("slicedI = ", slicedI)
	fmt.Println("i = ", i)
}
