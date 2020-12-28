/*
=====================defer keyword=============
#defer works on LIFO
#postpones execution of function util surrounding function returns
#widely used for output and input function and saves you from clos a open file
#deferred function executed in LIFO order.


*/

package main

import (
	"fmt"
)

func d1() {
	for i := 4; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}
func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}
func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}
func defer1() {
	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
}
