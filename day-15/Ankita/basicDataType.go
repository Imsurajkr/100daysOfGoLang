/*
====================BASIC GO DATA TYPE============================
Numeric Data type
Go has a native support for integers,floating-point numbers and complex numbers.
=====================Signed and unsigned integer================
Signed Integer(-127 to 127)          unsigned integer(0 to 255)
		1.int8                              1.uint8
		2.int16                             2.uint16
		3.int32                             3.uint32
		4.int64                             4.uint64

 you get to have seven binary digits to store your number because the eighth bit is used for keeping the
 sign of the integer. The same rule applies to the other sizes of unsigned integers.
(number at last of each type represents number of bits used)
NOTE:============= the size of these types changes depending on the architecture.

===============Floating point============================
Go supports only two types of floating-point numbers: float32 and float64.
float32=======provides about 6 decimal digits of precision
float64=========gives you 15 digits of precision.

===============Complex number================
a+ib
There are two complex number types
1.complex64=========uses two float32 one for real part and other for imaginary part.
2.complex128========uses two float64 one for real part and other for imaginary part.

*/

package main

import "fmt"

func main() {
	// i=9
	c1 := 23 + 4i
	// s := "ankita"
	// c9 = 23 +4*i
	c2 := complex(5, 7)
	//here we can not use Println,print instead of printf because Println call has possible formatting directive %T
	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type Of c2: %T", c2)
	fmt.Println(c1)
	fmt.Println(c2)
	var c3 complex64 = complex64(c1 + c2)
	fmt.Printf("Type Of %T", c3)
	fmt.Println(c3)
	c4 := c3 - c3
	fmt.Println(c4)
	x := 22
	y := 7
	div := x % y
	//%i can not be used with printf because Printf format %i has unknown verb i
	fmt.Printf("Type of x is %T\n", x)
	fmt.Println(div)
	fmt.Printf("Type of x is %T", div)
	var m, n float64
	m = 1.223
	fmt.Println("m, n:", m, n)
	y1 := 4 / 2.3
	// y1 = 22 / 7
	fmt.Printf("Type of divFloat: %T\n", y1)
	fmt.Println("y1:", y1)
	divFloat := float32(x) / float32(y)
	fmt.Println("divFloat", divFloat)
	fmt.Printf("Type of divFloat: %T\n", divFloat)
}
