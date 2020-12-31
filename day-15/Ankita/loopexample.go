package main

import (
	"fmt"
)

func spaces(n int) int {
	if n == 0 {
		return 0
	}
	fmt.Print(" ")
	spaces(n - 1)
	return 0
}

func stars(n int) int {
	if n == 0 {
		return 0
	}
	fmt.Print("*")
	stars(n - 1)
	return 0
}
func main() {
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}
		if i == 95 {
			break
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
	i := 10
	for {
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}
	fmt.Println()
	i = 0
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}
		fmt.Print(i, " ")
		i++
	}
	fmt.Println("how to iterate over an array using for loop")
	anArray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anArray {
		fmt.Println("index:", i, "value: ", value)
	}
	fmt.Println("Print a Triangle using only single for Loop")
	var count = 0

	for i := 0; i < 16; i++ {
		if count == 0 {
			spaces(16 - (i + 1))
			count = 1
		}
		if count == 1 {
			stars(2*i + 1)
			count = 0
		}
		if count == 0 {
			spaces(16 - (i + 1))
		}
		fmt.Println("")

	}

}
