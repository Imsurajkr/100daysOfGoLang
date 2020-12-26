package main

import (
	"fmt"
)

func a() {
	fmt.Println("Inside a()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()!")
		}
	}()
	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited!")
	fmt.Println("Exiting a()")
}

// The good thing is that panicRecover.go ended according to our will
// without panicking because the anonymous function used in defer took control of the
// situation. Also note that function b() knows nothing about function a() ; however,
// function a() contains Go code that handles the panic condition of function b() .
func b() {
	fmt.Println("Inside b()")
	panic("Panic in b()!")
	fmt.Println("Exiting b()")
}

// Change G to main
func g() {
	a()
	fmt.Println("main() ended!")
}
