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
The init() function is a private function by design, which means that it
cannot be called from outside the package in which it is contained.

feature is called Release candidate
bugs is called hotflix
Major changes is called release
we achieve this by a concept in git called Taging
using git tag we can specify the version in the code such as
git tag v0.1.0
git push origin v0.1.0
We have created a new tag successfully


*/
package main

import (
	"apackage"
	"fmt"
	"time"
)

func init() {
	fmt.Println("init() manyInit")
	fmt.Println(time.Now())
}
func main() {
	apackage.A()
	apackage.B()
	fmt.Println("Hello calling from main package")
}
