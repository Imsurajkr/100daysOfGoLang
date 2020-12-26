package main

import (
	"fmt"
)

func d7() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

// Replace d8 with main to execute the programme

func d8() {
	d7()
}

// Due to the parameter of the anonymous function, each time the anonymous
// function is deferred, it gets and therefore uses the current value of i . As a result, each
// execution of the anonymous function has a different value to process, hence the generated
// output.
