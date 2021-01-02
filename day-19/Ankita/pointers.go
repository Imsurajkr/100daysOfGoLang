/*
A pointer is a variable that stores the memory address of another variable.
======Declaring a pointer================================
var p *int;
No arithmetic operators are allowed in pointers in Go but they are allowed in other programming languages.
*/

package main

import "fmt"

func main() {
	var x = 100
	var p = &x
	p1 := &p
	fmt.Println(p)
	// You can use the * operator on a pointer to access the value stored in the variable that the pointer points to.
	// This is called dereferencing or indirecting -
	fmt.Println(*p)
	//*p1--
	fmt.Println(p1)

	// You can also create a pointer using the built-in new() function.
	//The new() function takes a type as an argument,
	// allocates enough memory to accommodate a value of that type, and returns a pointer to it.
	ptr := new(int)
	fmt.Println(ptr)
}
