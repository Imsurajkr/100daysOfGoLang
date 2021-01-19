package main

import (
	"fmt"
)

func main() {
	var myint interface{} = 123
	k, ok := myint.(int)
	if ok {
		fmt.Printf("convrted\n")
	}
	fmt.Println(k)
	l, ok := myint.(float64)
	if ok {
		fmt.Printf("convrted to float64")
	}
	fmt.Println(l)
	//without ok it return panic
	i := myint.(int)
	fmt.Println("No checking:", i)
	j := myint.(bool)
	fmt.Println(j)
}
