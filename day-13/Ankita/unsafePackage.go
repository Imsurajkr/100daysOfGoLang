/*
=============What is Unsafe code?================================
unsafe code is th code in go program that bypasses safety and security of yur program.mostly
it is related to pointers.using it is dangerous for your program.

=============

*/

package main

import (
	"fmt"
	"unsafe"
)

// func main() {
// 	var value int64 = 345
// 	var P1 = &value
// 	var P2 = (*int32)(unsafe.Pointer(P1))
// 	fmt.Println("P1 ", *P1)
// 	fmt.Println("P2 ", *P2)
// 	*P1 = 5434123412312431212
// 	fmt.Println(value)
// 	fmt.Println("P2 ", *P2)
// 	*P2 = 54341234
// 	fmt.Println(value)
// 	fmt.Println("*P2: ", *P2)
// }

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
	fmt.Println()
	pointer = (*int)(unsafe.Pointer(memoryAddress))
	fmt.Print("One more: ", *pointer, " ")
	memoryAddress = uintptr(unsafe.Pointer(pointer)) +
		unsafe.Sizeof(array[0])
	fmt.Println()
}
