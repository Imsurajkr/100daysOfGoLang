package main

// This is leared from the codes singh Go
import "fmt"

func main() {
	// Loops
	// For Loop
	for i := 1; i <= 10; i++ {
		for j := 2; j <= 2; j++ {
			fmt.Println("2 * ", i, "=", i*j)
		}
		//fmt.Println(i)
	}
	for x := 10; x > 5; x-- {
		fmt.Println("lets decrease", x)
	}
}

// This program print the tables of 2
