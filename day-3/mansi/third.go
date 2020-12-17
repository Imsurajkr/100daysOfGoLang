// //:=  the short assignment statement.
// The short assignment
// statement can be used in place of a var declaration with an implicit type.
// the var keyword is mostly used
// for declaring global variables in Go programs, as well as for declaring
// variables without an initial value.
// This means that the short assignment
// statement cannot be used outside of a function because it is not available
// there.
// The := operator works as follows:
// m := 123 is same as var m = 123 
package main
import "fmt"
func main(){
	num1 :=5
	num2 :=7
	result:=num1*num2
	fmt.Println("multiplication of num1 and num2 is ",result)
	// this is example of statictype variable declaration which tells the compiler that there is one variable 
	//available with the given type and name
	var x float64
	x=50.
	// Integer
	var2:=10

	//boolean
	var3:=true

	//string type 
	var4:="hello"

	//complex 64 and comlex128
	//is a basic data type


	fmt.Println(x)
	//we can check type of any variable using %T format specifier with printf
	   fmt.Printf("x is of TypeOf %T \n",x)
	   fmt.Printf("var2=%T\n",var2)
	   fmt.Printf("var3=%T\n",var3)
	   fmt.Printf("var4=%T\n",var4)
}


// output
// multiplication of num1 and num2 is  35
// 50
// x is of TypeOf float64
// var2=int
// var3=bool
// var4=string

