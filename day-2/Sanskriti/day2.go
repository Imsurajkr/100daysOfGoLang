//How to write comments in go?
// This is a single line comment.package Sanskriti
/*
	This is a multiple lines of comment.
	We use backslash and asterisk.
*/
// How to declare variabls in go?
// Different type of data type in Go and different type of ways to find out the data type of variables.

package main

// Here we import reflect along with fmt in order to see use that for finding the data type
import (
	"fmt"
	"reflect"
)

//main function

func main() {
	//How to declare variables?
	// First Method

	num := 20
	fmt.Println(num)

	// Second way
	/*
	  use keywors var then name of your variable and data type of that variable
	*/
	var a int = 10
	fmt.Println(a)

	// Now let's see different types of variables and methods to find there data type

	// string type
	var1 := "hello world"

	// integer
	var2 := 10

	// float
	var3 := 1.55

	// boolean
	var4 := true

	// shorthand string array declaration
	var5 := []string{"foo", "bar", "baz"}

	// map is reference datatype of hash table
	// golang Maps is a collection of unordered pairs of key-value.
	var6 := map[int]string{100: "Ana", 101: "Lisa", 102: "Rob"}

	// complex64 and complex128
	/* complex64 is complex numbers which have float32 real and imaginary parts.
	and complex128 is complex numbers which have float64 real and imaginary parts,
	golaang also have built-in function complex is used to construct a complex
	number with real and imaginary parts. */
	var7 := complex(9, 15)

	// using %T format specifier to
	// determine the datatype of the variables

	fmt.Println("Using Percent T with Printf")
	fmt.Println()

	fmt.Printf("var1 = %T\n", var1)
	fmt.Printf("var2 = %T\n", var2)
	fmt.Printf("var3 = %T\n", var3)
	fmt.Printf("var4 = %T\n", var4)
	fmt.Printf("var5 = %T\n", var5)
	fmt.Printf("var6 = %T\n", var6)
	fmt.Printf("var7 = %T\n", var7)

	// using TypeOf() method of reflect package
	// to determine the datatype of the variables
	fmt.Println()
	fmt.Println("Using reflect.TypeOf Function")
	fmt.Println()

	fmt.Println("var1 = ", reflect.TypeOf(var1))
	fmt.Println("var2 = ", reflect.TypeOf(var2))
	fmt.Println("var3 = ", reflect.TypeOf(var3))
	fmt.Println("var4 = ", reflect.TypeOf(var4))
	fmt.Println("var5 = ", reflect.TypeOf(var5))
	fmt.Println("var6 = ", reflect.TypeOf(var6))
	fmt.Println("var7 = ", reflect.TypeOf(var7))

	// using ValueOf() method of reflect package
	// to determine the value of the variable
	// Kind() method returns the datatype of the
	// value fetched by the ValueOf() method
	fmt.Println()
	fmt.Println("Using reflect.ValueOf.Kind() Function")
	fmt.Println()

	fmt.Println("var1 = ", reflect.ValueOf(var1).Kind())
	fmt.Println("var2 = ", reflect.ValueOf(var2).Kind())
	fmt.Println("var3 = ", reflect.ValueOf(var3).Kind())
	fmt.Println("var4 = ", reflect.ValueOf(var4).Kind())
	fmt.Println("var5 = ", reflect.ValueOf(var5).Kind())
	fmt.Println("var6 = ", reflect.ValueOf(var6).Kind())
	fmt.Println("var7 = ", reflect.ValueOf(var7).Kind())

}
