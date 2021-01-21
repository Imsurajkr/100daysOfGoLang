package main

import (
	"fmt"
)

type first struct{}

func (a first) F() {
	a.shared()
}

func (a first) shared() {
	fmt.Println("This is shared() from first")
}

type second struct {
	first
}

func (b second) shared() {
	fmt.Println("This is shared() from second")
}

type a struct {
	xx int
	yy int
}

type b struct {
	zz string
	ww int
}

type c struct {
	A a
	B b
}

func (A a) A() {
	fmt.Println("function A() for A.")
}
func (B b) A() {
	fmt.Println("function A() for B.")
}
func main() {
	first{}.F()
	second{}.shared()
	j := second{}
	k := j.first
	k.F()
	var i c
	i.A.A()
	i.B.A()
}
