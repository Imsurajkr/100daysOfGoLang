/*
   Go language provides a special feature known as an anonymous function.

   An anonymous function is a function which doesn’t contain any name.
   It is useful when you want to create an inline function.

   An anonymous function is also known as function literal.
*/

// How to create an anonymous function?
package main

import "fmt"

func main() {

	/* Anonymous function
	   Basic syntax for anonymous function is
	   Func(parameter_list)(return_type)
	   Use return statment is return_type is given
	*/

	func() {

		fmt.Println("Printing my output using anonymous function. ")
	}()

}

/*
   Important points about anonymous function.
   1. You are allowed to assign an anonymous function to a variable.
	  And when we do so the type of the variable is of function type.

   2. You can also pass arguments in the anonymous function.

   3. You can also pass an anonymous function as an argument into other function.

   4. You can also return an anonymous function from another function.
*/
