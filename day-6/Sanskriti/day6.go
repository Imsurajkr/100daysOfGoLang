package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func main() {

	//creating struct specifying field names
	emp1 := Employee{
		firstName: "Sanskriti",
		age:       25,
		salary:    500,
		lastName:  "Raghav",
	}

	//creating struct without specifying field names
	emp2 := Employee{"Ankita", "Mansi", 29, 800}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)
}
