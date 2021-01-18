package main

import (
	"fmt"
	"os"
	"strconv"
)

func star(l int) {
	if l < 0 {
		return
	}
	star(l - 1)
	fmt.Print("* ")
}

func holow(l int) {
	if l == 0 || l == 1 {
		return
	}
	holow(l - 1)
	fmt.Print("  ")
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please Provide height of pyramid")
		return
	}
	n, _ := strconv.Atoi(arguments[1])
	fmt.Println("Half Pyramid")
	for i := 0; i < n; i++ {
		star(i)
		fmt.Println("")
	}

	fmt.Println("Inverted Half Pyramid")
	for i := n - 1; i > -1; i-- {
		star(i)
		fmt.Println("")
	}

	fmt.Println("Hollow Half Pyramid")
	for i := n - 1; i > -1; i-- {
		if i == 0 || i == n-1 {
			star(i)
		} else {
			fmt.Print("* ")
			holow(i)
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}
