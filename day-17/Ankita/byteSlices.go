/*
============Multi-dimensional slices=============
Slices can have many dimensions just like arrays. The next statement creates a slice with
==============two dimensions:
s1 := make([][]int, 4)
If you find yourself using slices with many dimensions all the time, you
might need to reconsider your approach and choose a simpler design that
does not require multi-dimensional slices.
*/

package main

import (
	"fmt"
)

func main() {
	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	copy(a6, a4)
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	fmt.Println()
	b6 := []int{-10, 1, 2, 3, 4, 5}
	b4 := []int{-1, -2, -3, -4}
	fmt.Println("b6:", b6)
	fmt.Println("b4:", b4)
	copy(b4, b6)
	fmt.Println("b6:", b6)
	fmt.Println("b4:", b4)
	fmt.Println()
	array4 := [4]int{4, -4, 4, -4}
	s6 := []int{1, 1, -1, -1, 5, -5}
	copy(s6, array4[0:])
	fmt.Println("array4:", array4[0:])
	fmt.Println("s6:", s6)
	fmt.Println()
	array5 := [5]int{5, -5, 5, -5, 5}
	s7 := []int{7, 7, -7, -7, 7, -7, 7}
	copy(array5[0:], s7)
	fmt.Println("array5:", array5)
	fmt.Println("s7:", s7)
}
