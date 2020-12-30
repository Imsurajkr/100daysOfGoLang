/*
if you want to know the number of iterations that are going to be executed in advance, you cannot
use the range keyword.
The range keyword also works with Go maps, which makes it pretty
handy and my preferred way of iteration.
One of the biggest problems with arrays is out-of-bounds errors, which means trying to
access an element that does not exist. This is like trying to access the sixth element of an
array with only five elements. The Go compiler considers compiler issues that can be
detected as compiler errors because this helps the development workflow. Therefore, the
Go compiler can detect out-of-bounds array access errors:
./a.go:10: invalid array index -1 (index must be non-negative)
./a.go:10: invalid array index 20 (out of bounds for 2-element array)
*/

package main

import (
	"fmt"
)

func main() {
	anArray := [4]int{1, 2, 4, -4}
	twoD := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14,
		15, 16}}
	threeD := [2][2][2]int{{{1, 0}, {-2, 4}}, {{5, -1}, {7, 0}}}
	fmt.Println("The length of", anArray, "is", len(anArray))
	fmt.Println("The first element of", twoD, "is", twoD[0][0])
	fmt.Println("The length of", threeD, "is", len(threeD))
	for i := 0; i < len(threeD); i++ {
		v := threeD[i]
		for j := 0; j < len(v); j++ {
			m := v[j]
			for k := 0; k < len(m); k++ {
				fmt.Print(m[k], " ")
			}
		}
		fmt.Println()
	}
	for _, v := range threeD {
		for _, m := range v {
			for _, s := range m {
				fmt.Print(s, " ")
			}
		}
		fmt.Println()
	}
}
