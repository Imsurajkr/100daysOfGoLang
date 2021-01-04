package main

import (
	"fmt"
)

// The d2() function also contains a for loop and a defer statement, which
// will also be executed three times. However, this time, the defer keyword is applied to an
// anonymous function instead of a single fmt.Print()

func d4() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}

func d10() {
	d4()
}

// After the for loop has ended, the value of i is 0 , because it is that value of i that made the
// for loop terminate. However, the tricky point here is that the deferred anonymous function
// is evaluated after the for loop ends because it has no parameters, which means that it is
// evaluated three times for an i value of 0 , hence the generated output. This kind of
// confusing code is what might lead to the creation of nasty bugs in your projects, so try to
// avoid it.
