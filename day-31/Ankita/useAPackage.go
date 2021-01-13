/*
==============Private variables and functions==================
 private variables and functions from public ones is that private ones can
be strictly used and called internally in a package. Controlling which functions, constants,
and variables are public or not is also known as encapsulation.
Go follows a simple rule that states that functions, variables, types, and so forth that begin
with an uppercase letter are public, whereas functions, variables, types, and so on that
begin with a lowercase letter are private.
================INIT function=======================
Every Go package can optionally have a private function named init() that is
automatically executed at the beginning of the execution time.

*/
package main

import (
	"apackage"
	"fmt"
)

func main() {
	apackage.A()
	apackage.B()
	fmt.Println("Hello calling from main package")
}
