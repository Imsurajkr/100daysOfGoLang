package main

import (
	"fmt"
	"os"
	"strconv"
)

var count = 0

func spaces(n int) int {
	if n == 0 {
		return 0
	}
	fmt.Print(" ")
	spaces(n - 1)
	return 0
}
func holow(l int) {
	if l == 1 || l == 2 {
		return
	}
	holow(l - 1)
	fmt.Print("  ")
}
func stars(n int) int {
	if n == 0 {
		return 0
	}
	fmt.Print("* ")
	stars(n - 1)
	return 0
}
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please Provide height of pyramid")
		return
	}
	n, _ := strconv.Atoi(arguments[1])
	fmt.Println("Inverted Full Pyramid")
	for i := n; i > 0; i-- {
		spaces(count)
		stars(i)
		fmt.Println("")
		count = count + 1
	}
	fmt.Println("Full Pyramid")
	count = n - 1
	for i := 1; i < n+1; i++ {
		spaces(count)
		stars(i)
		fmt.Println("")
		count = count - 1
	}
	fmt.Println("Hollow Full Pyramid")
	count = n - 1
	for i := 1; i < n+1; i++ {
		spaces(count)
		if i == 1 || i == n {
			stars(i)
		} else {
			fmt.Print("* ")
			holow(i)
			fmt.Print("* ")
		}
		fmt.Println("")
		count = count - 1
	}
}
