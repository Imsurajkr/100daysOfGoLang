package main

import (
	"fmt"
	"os"
)

//chande defer to main
func defr() {
	if len(os.Args) == 1 {
		panic("Not enough arguments!")
	}
	fmt.Println("Thanks for the argument(s)!")
}

// Just Panic
// the use of the panic() and recover() pair is
// much more practical and professional than just using panic() .
