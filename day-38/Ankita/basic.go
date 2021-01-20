package main

import (
	"fmt"
	"reflect"
)

type s1 struct {
	A int
	B string
}
type s2 struct {
	C int
	D string
}

func main() {

	x := 100
	xRefl := reflect.ValueOf(&x).Elem()
	// xType := xRefl.Type()
	fmt.Printf("The type of x is %s.\n", xRefl)
	ref := reflect.TypeOf(&x)
	fmt.Print(ref)
	ans := s1{A: 10, B: "Ankita"}
	ques := s2{C: 67, D: "yadav"}
	fmt.Println(reflect.TypeOf(ans), reflect.ValueOf(ans))
	fmt.Println(reflect.TypeOf(ques), reflect.ValueOf(ques))
}
