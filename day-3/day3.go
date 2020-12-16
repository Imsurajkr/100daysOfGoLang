package main

import "fmt"

func main() {

	age := 10
	// To find person is able to vote or not
	if age>=18 {
		fmt.Println("you are able to vote")
	} else {
		fmt.Println("you are not able to vote")
	}
}