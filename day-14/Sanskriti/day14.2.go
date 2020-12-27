// Slice using array
/*
   For creating a slice from the given array first you need
   to specify the lower and upper bound, which means slice can
   take elements from the array starting from the lower bound
   to the upper bound. It does not include the elements above from the upper bound.
*/

package main

import "fmt"

func main() {

	// Creating an array
	arr := [4]string{"Hello", "How", "you", "Doin?"}

	// Creating slices from the given array
	var my_slice_1 = arr[1:2]
	my_slice_2 := arr[0:]
	my_slice_3 := arr[:2]
	my_slice_4 := arr[:]

	// Display the result
	fmt.Println("My Array: ", arr)
	fmt.Println("My Slice 1: ", my_slice_1)
	fmt.Println("My Slice 2: ", my_slice_2)
	fmt.Println("My Slice 3: ", my_slice_3)
	fmt.Println("My Slice 4: ", my_slice_4)
}

// The default value of the lower bound is 0
// and the default value of the upper bound
// is the total number of the elements present in the given array.
