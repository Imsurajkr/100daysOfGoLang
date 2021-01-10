package main

import (
	"fmt"
	"unsafe"
)

/*
   Here we storing the address of array[0] to a variable pointer.
   Then we convert the integer value of pointer into unsafe.Pointer() and then to
   uintptr() and store the result into a variable called memoryAddress.

	unsafe.Sizeof(array[0]) will give us to raech the next element of array,
	as the size occupied by each elements of the array is equal.

	And we use for loop to find the memory address of each element by adding the unsafe.Sizeof(array[0])
	to memoryAddress variable in each iteration.
*/

func main() {
	array := [...]int{0, 1, -2, 3, 4}
	pointer := &array[0]
	fmt.Print(*pointer, " ")
	memoryAddress := uintptr(unsafe.Pointer(pointer)) +
		unsafe.Sizeof(array[0])

	for i := 0; i < len(array)-1; i++ {
		pointer = (*int)(unsafe.Pointer(memoryAddress))
		fmt.Print(*pointer, " ")
		memoryAddress = uintptr(unsafe.Pointer(pointer)) +
			unsafe.Sizeof(array[0])
	}

	// Here we are trying to access an element of an array which doesn't access,
	// And because we are using unsafe package the Go compiler can not catch such a logical error.
	// Hence it returns something inaccurate.
	fmt.Println()
	pointer = (*int)(unsafe.Pointer(memoryAddress))
	fmt.Print("One more: ", *pointer, " ")
	memoryAddress = uintptr(unsafe.Pointer(pointer)) +
		unsafe.Sizeof(array[0])
	fmt.Println()
}
