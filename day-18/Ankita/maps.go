/*
=================Maps=================
A map is an unordered collection of key value pairs
The Keys are unique in map while the values may not be.
The map data structure is used for fast lookups, retrieval, and deletion of data
 based on keys. It is one of the most used data structures in computer science.
Maps are reference types.
===================When you should use a map?==========
1.Maps are more versatile than both slices and arrays but this flexibility comes at a cost:
 the extra processing power required for the implementation of a Go map. However, built-in Go
structures are very fast, so do not hesitate to use a Go map when you need to. What you
should remember is that Go maps are very convenient and can store many different kinds
of data, while being both easy to understand and fast to work with.
*/

package main

import "fmt"

func main() {
	var m map[string]int
	// m["one"] = 100
	fmt.Println(m)

	var s = make(map[string]int)
	s["one"] = 100
	fmt.Println(s)

	var m1 = map[int]int{
		1:  90,
		78: 678,
		43: 500,
		// last trailing comma is necessary, otherwise, you’ll get a compiler error.
	}
	fmt.Println(m1)
	var m2 = map[int]int{}
	//If you try to add a key that already exists in the map, then it will simply be overridden by the new value.
	m2[90] = 100
	m2[80] = 90
	m2[70] = 80
	m2[60] = 70
	m2[50] = 60
	m2[50] = 65
	fmt.Println(m2)
	fmt.Println("value of m2 at index 60", m2[60])
	//check for the existence of a key in a map
	name, ok := m2[80]
	fmt.Println(ok, name)
	delete(m2, 50)
	fmt.Println(m2)
	Ankita := m2
	fmt.Println(Ankita)
	Ankita[80] = 00
	fmt.Println(Ankita)
	fmt.Println(m2)
	// A map is an unordered collection and therefore the iteration order
	//of a map is not guaranteed to be the same every time you iterate over it.
	for index, val := range m2 {
		fmt.Println(index, "=", val)
	}
}
