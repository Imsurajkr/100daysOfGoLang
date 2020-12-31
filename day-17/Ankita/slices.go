/*
1.A Slice is a segment of an array. Slices build on arrays and provide more power,
flexibility, and convenience compared to arrays.
2.Just like arrays, Slices are indexable and have a length.But unlike arrays,
they can be resized.
3.A Slice is just a reference to an underlying array.
4.slices are passed by reference to functions, which means that what is actually passed is
the memory address of the slice variable, any modifications you make to a slice inside a
function will not get lost after the function exits.
5.passing a big slice to a function is significantly faster than passing an array with the
same number of elements because Go will not have to make a copy of the slice; it will just
pass the memory address of the slice variable.
=======================A slice consists of three things===============
1.A pointer (reference) to an underlying array.
2.The length of the segment of the array that the slice contains.
3.The capacity (the maximum size up to which the segment can grow).
*/

package main

import "fmt"

func main() {
	var s = []int{1, 2, 3, 4}
	var array = [5]string{"ankita", "mansi", "suraj", "anurag", "Sanskriti"}
	//create a slice using an array
	var g = array[2:5]
	slice1 := array[1:4]
	slice2 := array[:3]
	slice3 := array[2:]
	slice4 := array[:]
	//creating slice from another slice
	slice5 := slice1[:1]
	//modification of slice
	slice5[0] = "anshul"
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(g)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
}
