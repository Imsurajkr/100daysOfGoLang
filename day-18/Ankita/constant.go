/*
 constants are variables that cannot change their values.
 constants are defined with the help of the const keyword.
 the value of a constant variable is defined at compile time not at run time.
NOTE:-GO is a statically typed language.
Go doesn’t provide implicit type conversions and it requires us to do explicit type casting
 whenever we mix variables of multiple types in an operation.
*/

package main

import (
	"fmt"
)

// type is a way of defining a new named type that uses the same
// underlying type as an existing type. This is mainly used for differentiating
// between different types that might use the same kind of data.
type Digit int
type Power2 int

const PI = 3.1415926
const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(PI)
	//The constant generator iota is used for declaring a sequence of related values that use
	//incrementing numbers without the need to explicitly type each one of them.
	const (
		Zero Digit = iota
		One
		Two
		Three
		Four
	)
	fmt.Println(One)
	fmt.Println(Three)
	//fmt.Println(Digit)
	fmt.Println(Two)
	const (
		p2_0 Power2 = 1 << iota
		_
		p2_2
		_
		p2_4
		_
		p2_6
	)
	fmt.Println("2^0:", p2_0)
	fmt.Println("2^2:", p2_2)
	fmt.Println("2^4:", p2_4)
	fmt.Println("2^6:", p2_6)
}
