// Using make() function
/*
  You can also create a slice using the make() function
  which is provided by the go library. This function takes
  three parameters, i.e, type, length, and capacity.
  It assigns an underlying array with a size that is equal
  to the given capacity and returns  a slice which refers to the
  underlying array. Generally, make() function is used to create an
  empty slice. Here, empty slices are those slices that contain an empty array reference.
*/

package main

import "fmt"

func main() {

	// Creating an array and Using make() function
	var my_slice_1 = make([]int, 4, 7)
	fmt.Printf("Slice 1 = %v, \nlength = %d, \ncapacity = %d\n",
		my_slice_1, len(my_slice_1), cap(my_slice_1))

	// Creating another array and Using make function
	var my_slice_2 = make([]int, 7)
	fmt.Printf("Slice 2 = %v, \nlength = %d, \ncapacity = %d\n",
		my_slice_2, len(my_slice_2), cap(my_slice_2))

}
