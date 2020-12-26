/*
slice is a segment of arrays.slice build on arrays and provide more flexibility and convinience
compared to arrays.
slice have indexes as well as length and they can be resized.
var ankita []int;
 var arr =[1,2,3,4,5,6,7,8,9]
 var slice[low:high]
slices are declared like arrays only except the fact we do not have to specify the size in the brackets.
*/

package main

import (
	"fmt"
	"runtime"
)

type slice struct {
	ankita, mansi int
}

func main() {
	var N = 400000
	var data []slice
	for i := 0; i < N; i++ {
		value := i
		data = append(data, slice{value, value})
	}
	fmt.Print(data)
	runtime.GC()
	_ = data[0]
}
