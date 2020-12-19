// Write a Go program that finds the average value of floating-point numbers that
// are given as command-line arguments.
// @author :- Ankita & Suraj
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var average = 0.0
	var sum = 0.0
	if len(os.Args) == 1 {
		fmt.Println("Please prvoide some input")
	} else {
		for i := 1; i < len(os.Args); i++ {
			fmt.Println("For Executted ------")
			n, _ := strconv.ParseFloat(os.Args[i], 64)
			fmt.Println("Value of n is ", n)
			sum = n + sum
			fmt.Println("Value of sum is ", sum)
		}
		average = sum / float64(len(os.Args)-1)
		fmt.Println("Value of average is ", average)

	}
}
