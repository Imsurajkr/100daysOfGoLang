package main

import "fmt"

func main() {
	var i [5]int = [5]int{1, 2, 3, 4, 5}
	var slicedI []int = i[:]
	fmt.Println(slicedI)
}
