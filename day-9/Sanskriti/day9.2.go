/* Golang program to the use of Blank identifier.
   In place of div identifier we just use the _ (underscore).
   It allows the compiler to ignore declared and not used error for that particular variable.
*/

package main

import "fmt"

// Main function
func main() {

	// calling the function
	// function returns two values which are
	// assigned to mul and blank identifier
	mul, _ := mul_div(105, 7)

	// only using the mul variable
	fmt.Println("105 x 7 = ", mul)
}

// function returning two
// values of integer type
func mul_div(n1 int, n2 int) (int, int) {

	// returning the values
	return n1 * n2, n1 / n2
}

/*
   1.You can use multiple Blank Identifier in the same program.
     So you can say a Golang program can have multiple variables using
     the same identifier name which is the blank identifier.

   2.There are many cases that arise the requirement of assignment of
     values just to complete the syntax even knowing that the values will
     not be going to be used in the program anywhere. Like a function returning
     multiple values. Mostly blank identifier is used in such cases.

   3.You can use any value of any type with the Blank Identifier.
*/
