/*
===================Arrays============================
arrays is one of the most popular topic of data structures because of two reasons:
1.they are simple and easy to understand
2.they are very versatile and can store many different kinds of data.
=========Declaration of an array============================
array := [8]int
array[0]=9
for i := 0; i <len(array);i++ {

}
anArray := [4]int{1, 2, 4, -4}
====================Multi-Dimensional array================
twoD := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12},
 {13, 14, 15, 16}}
threeD := [2][2][2]int{{{1, 0}, {-2, 4}}, {{5, -1}, {7, 0}}}
==================DISAADVANTAGES OF ARRAY===============.
1 array cannot change its size, which means that Go arrays are not dynamic.
2.when you pass an array to a function as a parameter,you actually pass a copy of the array, which means that any changes you make to an array
inside a function will be lost after the function exits.
3.passing a large array to a function can be pretty slow, mostly because Go has to create a copy of the array.
##NOTE:==Because of their disadvantages, arrays are rarely used in Go!
*/

package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 4, -4}
	//the range keyword allow you to bypass the use of the len() function in the for loop
	for i, val := range arr {
		fmt.Println("index is: ", i, " value is: ", val)
	}
	fmt.Print(len(arr))
}
