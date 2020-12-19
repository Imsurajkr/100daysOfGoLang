package main

import "fmt"

func main() {
	// This is for loop syntax
	for i := 1; i < 100; i++ {

		fmt.Println(i)
	}

	// Using continue and break keyword
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}
		if i == 95 {
			break
		}
		fmt.Print(i, " ")
	}

	// Using array and range keyword
	anarray := [10]string{"Hi", "I'm", "Learning", "Go", "Language", "So", "Far", "It", "is", "going"}
	for i, value := range anarray {
		fmt.Println(i, value)
	}

}
