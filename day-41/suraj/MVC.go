// Model

// The Model component corresponds to all the data-related
// logic that the user works with. This can represent either
// the data that is being transferred between the View and Controller
// components or any other business logic-related data.
// For example, a Customer object will retrieve the customer information
// from the database, manipulate it and update it data back to the
// database or use it to render data.

// View

// The View component is used for all the UI logic of the application.
// For example, the Customer view will include all the UI components such as text boxes,
// dropdowns, etc. that the final user interacts with.

// Controller

// Controllers act as an interface between Model and View components to process all the business logic
// and incoming requests, manipulate data using the Model component and interact with the Views to
// render the final output. For example, the Customer controller will handle all the interactions
// and inputs from the Customer View and update the database using the Customer Model.
// The same controller will be used to view the Customer data
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
