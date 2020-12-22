/*
	_(underscore) in Golang is known as the Blank Identifier.

	Identifiers are the user-defined name of the program components used for
	the identification purpose. Golang has a special feature to define and use
	the unused variable using Blank Identifier.

	The real use of Blank Identifier comes when a function returns multiple values,
	but we need only a few values and want to discard some values.

	Basically, it tells the compiler that this variable is not needed and ignored it without any error.

	It hides the variable’s values and makes the program readable. So whenever you will
	assign a value to Bank Identifier it becomes unusable.
*/

/*
   This program shows the compiler
   throws an error if a variable is
   declared but not used.
*/

package main

import "fmt"

// Main function
func main() {

	// calling the function
	// function returns two values which are
	// assigned to mul and div identifier.
	mul, div := mul_div(105, 7)

	// only using the mul variable
	// compiler will give an error
	fmt.Println("105 x 7 = ", mul)
}

// function returning two
// values of integer type
func mul_div(n1 int, n2 int) (int, int) {

	// returning the values
	return n1 * n2, n1 / n2
}

/*
   In the above program, the function mul_div is returning two values
   and we are storing both the values in mul and div identifier.
   But in the whole program, we are using only one variable i.e. mul.
   So compiler will throw an error div declared and not used.
*/
/*
	Let’s make the use of Blank identifier to correct the above program.
	Have a look on the program day9.2 for the same.

*/
