/*
  In Go language slice is more powerful, flexible, convenient
  than an array, and is a lightweight data structure. Slice is
  a variable-length sequence which stores elements of a similar type,
  you are not allowed to store different type of elements in the same slice.
  It is just like an array having an index value and length, but the size of the
  slice is resized they are not in fixed-size just like an array. Internally,
  slice and an array are connected with each other, a slice is a reference to
  an underlying array. It is allowed to store duplicate elements in the slice.

*/

/*
  Component of Slice:

  1. Pointer - The pointer is used to points to the first element of the array
	 that is accessible through the slice. Here, it is not necessary that the pointed
	 element is the first element of the array.

  2. Length - The length is the total number of elements present in the array.

  3. Capacity - The capacity represents the maximum size upto which it can expand.
*/

package main

import "fmt"

func main() {

	// Creating an array
	arr := [7]string{"This", "is", "day-13", "of",
		"GoLang", "Programming", "chanllenge"}

	// Displaying array
	fmt.Println("Array:", arr)

	// Creating a slice
	myslice := arr[1:6]

	// Displaying slice
	fmt.Println("Slice:", myslice)

	// Displaying length of the slice
	fmt.Printf("Length of the slice: %d", len(myslice))

	// Displaying the capacity of the slice
	fmt.Printf("\nCapacity of the slice: %d", cap(myslice))
}

/*
   A slice can be created and initialized using:
   1. Using Literal
   2. Using an array
   3. Using already existing slices
   4. Using make() function
*/
/*
   Ways to iterate over a slice:
   1. Using for loop
   2. Using range in for loop
   3. Using a blank identifier in for loop
*/

/*
   Important points about slice:
   1. You are allowed to create a nil slice that does not contain
	  any element in it. So the capacity and the length of this slice is 0.

   2. As slice is a reference type it can refer an underlying array.
	  So,if you made any changes in the slice, then it will also reflect in the array also.

   3. You can only use == operator to check whether the given slice is same or not.

	  If you try to compare two slices with the help of != operator then it will give you an error.
   4. Multi-dimensional slice are just like the multidimensional array, except that slice does not contain the size.

   5. You are allowed to sort the elements present in the slice.
	  The standard library of Go language provides the sort package ( sort() ) which contains different
	  types of sorting methods for sorting  the slice of ints, float64s, and strings. These functions always
	  sort the elements available is slice in ascending order.
*/
