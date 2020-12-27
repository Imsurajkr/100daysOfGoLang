// Using already Existing Slice
/*
  For creating a slice from the given slice first you need to
  specify the lower and upper bound, which means slice can take
  elements from the given slice starting from the lower bound to
  the upper bound. It does not include the elements above from the upper bound.
*/

package main

import "fmt"

func main() {

	// Creating s slice
	orignal_slice := []int{90, 60, 40, 50,
		34, 49, 30}

	// Creating slices from the given slice
	var my_slice_1 = orignal_slice[1:5]
	my_slice_2 := orignal_slice[0:]
	my_slice_3 := orignal_slice[:6]
	my_slice_4 := orignal_slice[:]
	my_slice_5 := my_slice_3[2:4]

	// Display the result
	fmt.Println("Original Slice:", orignal_slice)
	fmt.Println("New Slice 1:", my_slice_1)
	fmt.Println("New Slice 2:", my_slice_2)
	fmt.Println("New Slice 3:", my_slice_3)
	fmt.Println("New Slice 4:", my_slice_4)
	fmt.Println("New Slice 5:", my_slice_5)
}

// The default value of the lower bound is 0 and the default value of the upper bound
// is the total number of the elements present in the given slice.
