package main

import (
	"fmt"
)

func main() {
	emp3 := struct {
		firstName string
		lastName  string
		age       int
		salary    int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}
	fmt.Println("First Name:", emp3.firstName)
	fmt.Println("Last Name:", emp3.lastName)
	fmt.Println("Age:", emp3.age)
	fmt.Printf("Salary: $%d\n", emp3.salary)
	emp3.salary = 6500
	fmt.Printf("New Salary: $%d", emp3.salary)
	fmt.Println("Employee 3", emp3)
}
