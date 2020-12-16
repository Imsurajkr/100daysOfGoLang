package main

import "fmt"

func main() {

	num1 := 10
	// To find num1 is odd or even
	if num1%2 == 0 {
		fmt.Println("Given number is even")
	} else {
		fmt.Println("Given number is odd")
	}
}

// if we start writting else code block before the paranthesis of else block it gives us a syntax error.
// So also start else block on the same line of closing bracket of if block.
