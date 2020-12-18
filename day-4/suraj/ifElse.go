package main

import "fmt"

// func main() {
// 	num1 := 10
// 	num2 := 20
// 	// Check if num1 is greater than num2 or num2 is greater than num1
// 	if num2 > num1 {
// 		fmt.Println("Bigger")
// 	} else {
// 		fmt.Println("Smaller")
// 	}
// }

func main() {
	num1 := 30
	num2 := 10
	// Check if num1 is greater than num2 or num2 is greater than num1
	if num2 > num1 {
		fmt.Println("Bigger")
	} else if num2 == num1 {
		fmt.Println("Same")
	} else {
		fmt.Println("Smaller")
	}
}
