/*
constants are variables that cannot change their values.
constants are defined with the help of the const keyword.
the value of a constant variable is defined at compile time not at run time.
NOTE:-GO is a statically typed language.
Go doesn’t provide implicit type conversions and it requires us to do explicit type casting
 whenever we mix variables of multiple types in an operation.
============ go do not support implicit conversion?
Reason----The convenience of automatic conversion between numeric types in C is outweighed by
the confusion it causes. When is an expression unsigned? How big is the value? Does it overflow?
Is the result portable, independent of the machine on which it executes? It also complicates
the compiler; “the usual arithmetic conversions” are not easy to implement and inconsistent
across architectures. For reasons of portability, we decided to make things clear and
straightforward at the cost of some explicit conversions in the code.

====UNTYPED CONSTANTS
Any constant in golang, named or unnamed, is untyped unless given a type explicitly like var a = 1.00

===============Constants and Type inference: Default Type
Go supports type inference. That is, it can infer the type of a variable from the value that is used
to initialize it. So you can declare a variable with an initial value, but without any type information,
 and Go will automatically determine the type -
Constants	          Default Type
integers (10, 76)	      int
floats (3.14, 7.92)    	float64
complex numbers (3+5i) 	complex128
characters ('a', '♠')   	rune
booleans (true, false)  	bool
strings (“Hello”)	       string
Typed Constants
In Golang, Constants are typed when you explicitly specify the type in the declaration like this-

const typedInt int = 1  // Typed constant
With typed constants, you lose all the flexibility that comes with untyped constants like assigning
them to any variable of compatible type or mixing them in mathematical operations. So you should declare
a type for a constant only if it’s absolutely necessary. Otherwise, just declare constants without a type.
*/

package main

import (
	"fmt"
)

//type keyword is used to create type alias
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
	var harsh int64 = 123
	var ankita int64 = harsh
	fmt.Println(ankita)
	fmt.Println(v1)
	fmt.Println(PI)
	type RichString string
	var myString string = "Hello"
	//below line produces error because Given the strongly typed nature of Golang, you can’t assign a string
	// variable to a RichString variable-
	{ /*var myRichString RichString = myString*/
	}
	// But, you can assign an untyped string constant to a RichString variable
	//  because it is compatible with strings -
	const myUntypedString = "Hello"
	var myRichString RichString = myUntypedString
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
