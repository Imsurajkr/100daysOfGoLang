/*
   Here we discussing how our way of storing pointer will impact the performance of garbage collection.
   We will comparing the perfromance of garbage collector by storing pointers using Slice , maps , maps without pointers ,
   and splitting maps.
   While comparing the performamce we found out that slice collaborate with Go garbage collector much better while
   maps slow down the performance of garbage collector.
*/
package main

import (
	"runtime"
)

func main() {
	var N = 40000000
	myMap := make(map[int]int)
	for i := 0; i < N; i++ {
		value := int(i)
		myMap[value] = value
	}
	runtime.GC()
	_ = myMap[0]
}
