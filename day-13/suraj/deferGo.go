package main

// The function that is deferred in d1() is
// the fmt.Print() statement; as a result, when the d1() function is about to return, you get
// the three values of the i variable of the for loop in reverse order because deferred
// functions are executed in LIFO order.
import (
	"fmt"
)

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

func d2() {
	for i := 1; i < 10; i++ {
		defer fmt.Print(i, " ")
	}
}

// here the output should be 123456789 but its opposite its because of the defer keyword
// Its very important to understand that defer keyword follows the Last In Fist Out
// And we call it through fmt package

func d11() {
	d2()
}

// Go code implements a function named d1()
// with a for loop and a defer statement that will be executed 9 times for d2 and for d1 3 .
