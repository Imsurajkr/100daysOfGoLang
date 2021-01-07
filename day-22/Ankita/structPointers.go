/*
 C or C++ background, it is perfectly legal for a Go
function to return the memory address of a local variable.

Structures are very important in Go and are used extensively in real world programs because
they allow you to group as many values as you want and treat those values as a single entity


=================NEW KEYWORD =================
# Go supports the new keyword, which allows you to allocate new objects.
# new returns the memory address of the allocated object.
# new returns a pointer.
# The main difference between new and make is that variables created with make are properly initialized
without just zeroing the allocated memory space. Additionally, make can only be applied to maps, channels,
and slices, and does not return a memory address, which means that make does not return a pointer.
*/
package main

import (
	"fmt"
)

type myStructure struct {
	Name    string
	Surname string
	Height  int32
}

func createStruct(n, s string, h int32) *myStructure {
	if h > 300 {
		h = 0
	}
	return &myStructure{n, s, h}
}
func retStructure(n, s string, h int32) myStructure {
	if h > 300 {
		h = 0
	}
	return myStructure{n, s, h}
}
func main() {
	s1 := createStruct("Mihalis", "Tsoukalos", 123)
	s2 := retStructure("Mihalis", "Tsoukalos", 123)
	fmt.Println((*s1).Name)
	fmt.Println(s2.Name)
	fmt.Println(s1)
	fmt.Println(s2)
	s3 := new(myStructure)
	fmt.Println(s3)
	sP := new([]myStructure)
	fmt.Println(sP)
}

/*
   var a = 10
   var int* b = &a
   fmt.println(*b)
*/
