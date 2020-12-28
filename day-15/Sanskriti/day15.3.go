// Pointers to a struct

// Pointers in Go programming language or Golang is a variable
// which is used to store the memory address of another variable.

// Golang program to illustrate
// the pointer to struct
package main

import "fmt"

// defining a structure
type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {

	// passing the address of struct variable
	// emp8 is a pointer to the Employee struct
	emp8 := &Employee{"Sam", "Anderson", 55, 6000}

	// (*emp8).firstName is the syntax to access
	// the firstName field of the emp8 struct
	fmt.Println("First Name:", (*emp8).firstName)
	fmt.Println("Age:", (*emp8).age)
}
