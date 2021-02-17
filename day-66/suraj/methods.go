package main

import "fmt"

//declared the structure named emp
type emp struct {
	name    string
	address string
	age     int
}

//Declaring a function with receiver of the type emp
func (e emp) display() {
	fmt.Println(e.name)
}

func main() {
	//declaring a variable of type emp
	var empdata1 emp

	//Assign values to members
	empdata1.name = "Sura"
	empdata1.address = "A32 SK-Singh"
	empdata1.age = 24

	//declaring a variable of type emp and assign values to members
	empdata2 := emp{
		"Ravi", "Building-24, Delhi", 25}

	//Invoking the method using the receiver of the type emp
	// syntax is variable.methodname()
	empdata1.display()
	empdata2.display()
}
