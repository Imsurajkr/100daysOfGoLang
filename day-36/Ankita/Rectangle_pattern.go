package main

import (
	"fmt"
	"os"
	"strconv"
)

func stars(l int) {
	if l == 0 {
		return
	}
	stars(l - 1)
	fmt.Print("* ")
}

func hollow(l int) {
	if l == 0 || l == 2 {
		return
	}
	hollow(l - 1)
	fmt.Print("  ")
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please Provide length and breadth")
		return
	}
	fmt.Println("Solid Rectangle")
	l, _ := strconv.Atoi(arguments[2])
	b, _ := strconv.Atoi(arguments[1])
	for i := 0; i < l; i++ {
		stars(b)
		fmt.Println("")
	}
	fmt.Println("Hollow Rectangle")
	for i := 0; i < l; i++ {
		if i == 0 || i == l-1 {
			stars(b)
		} else {
			fmt.Print("* ")
			hollow(b)
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}
