/*
	What are type Methods?
	Type methods are go functions with some special reciever argument.
	implementation of function is exactly the same as the implementation of methods !
*/

package main

import (
	"fmt"
)

type twoPoint struct {
	X int64
	Y int64
}

func normal(a, b twoPoint) twoPoint {
	num := twoPoint{X: a.X + b.X, Y: a.Y + b.Y}
	return num
}

func (a twoPoint) typemethods(b twoPoint) twoPoint {
	num := twoPoint{X: a.X + b.X, Y: a.Y + b.Y}
	return num
}

func main() {
	a := twoPoint{X: 6, Y: 9}
	b := twoPoint{X: 66, Y: 99}
	fmt.Println(normal(a, b))
	fmt.Println(a.typemethods(b))
}
