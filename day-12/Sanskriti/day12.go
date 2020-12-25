/*
  Go language provides inbuilt support for basic constants
  and mathematical functions for complex numbers with the help
  of the cmplx package. You are allowed to find the cosine of the
  specified complex number with the help of the Cos() function provided
  by the math/cmplx package. So, you need to add a math/cmplx package in
  your program with the help of the import keyword to access the Cos() function.
*/

// Program to find the cosine value of the given complex number.

package main

import (
	"fmt"
	"math/cmplx"
)

// Main function
func main() {

	// Using Cos() function
	res_1 := cmplx.Cos(2 + 5i)
	res_2 := cmplx.Cos(-1 + 8i)
	res_3 := cmplx.Cos(-1 - 7i)

	// Displaying the result
	fmt.Println("Result 1:", res_1)
	fmt.Println("Result 2:", res_2)
	fmt.Println("Result 3:", res_3)

	// Another program to find cosine using cos() function
	// Declaring complex number

	cnumber_1 := complex(1, 2)
	cnumber_2 := complex(3, 6)

	// Finding cosine
	cvalue_1 := cmplx.Cos(cnumber_1)
	cvalue_2 := cmplx.Cos(cnumber_2)

	// Sum of two cosine values
	res := cvalue_1 + cvalue_2

	// Displaying results
	fmt.Println("Complex Number 1: ", cnumber_1)
	fmt.Println("Cosine 1: ", cvalue_1)

	fmt.Println("Complex Number 2: ", cnumber_2)
	fmt.Println("Cosine 2: ", cvalue_2)
	fmt.Println("Sum : ", res)
}
