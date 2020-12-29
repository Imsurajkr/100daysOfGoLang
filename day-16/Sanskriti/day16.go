/*
   Go language provides inbuilt support for bits to implement
   bit counting and manipulation functions for the predeclared
   unsigned integer types with the help of bits package.
   This package provides ReverseBytes() function which is used to
   find the reversed order of the value of a. To access ReverseBytes()
   function you need to add a math/bits package in your program with
   the help of the import keyword.
*/

// This function takes one parameter of uint type
// This function returns the value of a with its bits in reversed order.

// Golang program to illustrate
// bits.ReverseBytes() Function

package main

import (
	"fmt"
	"math/bits"
)

// Main function
func main() {

	// Finding the reverse order of a
	a := bits.ReverseBytes(7)
	fmt.Printf("Reverse order of %d: %b", 7, a)

	a1 := bits.ReverseBytes(3)
	fmt.Printf("ReverseBytes(%b) := %b\n", 3, a1)

	a2 := bits.ReverseBytes(11)
	fmt.Printf("ReverseBytes(%b) := %b\n", 11, a2)
}
