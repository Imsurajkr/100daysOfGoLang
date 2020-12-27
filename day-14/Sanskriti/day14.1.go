// Slice using literals
/*
   You can create a slice using the slice literal.
   The creation of slice literal is just like an array literal,
   but with one difference you are not allowed to specify
   the size of the slice in the square braces[].
*/

package main

import "fmt"

func main() {

	// Creating a slice
	var my_slice_1 = []string{"How", "you", "Doin?"}

	fmt.Println("My Slice 1:", my_slice_1)

	// Creating a slice
	//using shorthand declaration
	my_slice_2 := []int{12, 45, 67, 56, 43, 34, 45}
	fmt.Println("My Slice 2:", my_slice_2)
}

//  Always remember when you create a slice using a string literal,
//  then it first creates an array and after that return a slice reference to it.
