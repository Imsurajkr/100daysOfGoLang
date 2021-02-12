package main

import "fmt"

//declared the structure named emp
type emp struct {
	name    string
	address string
	age     int
}

//function which accepts variable of emp type and prints name property
func display(e emp) {
	fmt.Println(e.name)
}

func main() {
	// declares a variable, empdata1, of the type emp
	var empdata1 emp
	//assign values to members of empdata1
	empdata1.name = "John"
	empdata1.address = "Street-1, London"
	empdata1.age = 30

	//declares and assign values to variable empdata2 of type emp
	empdata2 := emp{"Raj", "Building-1, Paris", 25}

	//prints the member name of empdata1 and empdata2 using display function
	display(empdata1)
	display(empdata2)
}
