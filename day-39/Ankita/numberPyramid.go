package main

import (
	"fmt"
)

func triangle(i int) {
	if i == 0 {
		return
	}
	triangle(i - 1)
	fmt.Print(i, " ")
}

func Hollow(i, j int) {
	if i == 0 {
		return
	} else if i == 1 || i == j+1 {
		Hollow(i-1, j)
		fmt.Print(i, " ")
	} else {
		Hollow(i-1, j)
		fmt.Print("  ")
	}
}

func main() {
	fmt.Println("Half Number Pyramid")
	for i := 0; i < 5; i++ {
		triangle(i + 1)
		fmt.Println("")
	}
	fmt.Println("Inverted Half Number Pyramid")
	for j := 5; j > 0; j-- {
		triangle(j)
		fmt.Println("")
	}
	fmt.Println("Hollow Half Pyramid")
	for i := 0; i < 5; i++ {
		if i == 0 || i == 4 {
			triangle(i + 1)
		} else {
			Hollow(i+1, i)
		}

		fmt.Println("")
	}
}
