package main

import (
	"fmt"
	"os"
	"strconv"
)

func diffvalues(num int) (int, int, int, float64) {
	return num + 10, num - 5, num * 6, float64(num) / float64(9)
}

func namedfunc(n int, n1 int) (min, max int) {
	if n > n1 {
		max = n
		min = n1
	} else {
		max = n1
		min = n
	}
	return
}
func pointfunc(v *int) int {
	fmt.Println(*v**v, "*************************8")
	return *v * *v
}

func ret_pointers(v *int) *int {
	mul := *v * *v
	return &mul
}

func anotherfunc(v *int) func() int {
	mul := *v * *v
	return func() int {
		return pointfunc(&mul)
	}
}

func nestedfunc(f func(*int) int, v int) int {
	return f(&v)
}

func main() {
	argument := os.Args
	n, _ := strconv.Atoi(argument[1])
	n1, _ := strconv.Atoi(argument[2])
	if len(argument) == 1 {
		fmt.Println("Please enter two numbers")
	} else {
		fmt.Println(namedfunc(n, n1))
	}
	fmt.Println(nestedfunc(pointfunc, n), "!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println(pointfunc(&n))
	fmt.Println(anotherfunc(&n), "6666666666666666666666")
	fmt.Println(ret_pointers(&n))
	fmt.Println(diffvalues(100))
}
