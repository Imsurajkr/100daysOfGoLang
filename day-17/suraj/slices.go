package main

import "fmt"

func main() {
	var s = []int{1, 2, 3, 4}
	var fruits = [5]string{"apple", "mango", "banana", "Litchi", "apple"}
	//create a slice using an array
	var g = fruits[2:5]
	slice1 := fruits[1:4]
	slice2 := fruits[:3]
	slice3 := fruits[2:]
	slice4 := fruits[:]
	//creating slice from another slice
	slice5 := slice1[:1]
	//modification of slice
	slice5[0] = "grapes"
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(g)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
}
