/*Write a Go program that finds the sum of all command-line arguments that are
valid numbers*/

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	mainString := ""
	var sum = 0.0
	var average = 0.0

	if len(os.Args) == 1 {
		mainString = "Please provide one input at least"
	} else {
		for i := 1; i < len(os.Args); i++ {
			n, _ := strconv.ParseFloat(os.Args[i], 64)
			sum = sum + n

		}
		average = sum / float64(len(os.Args))
		fmt.Println("Sum of command line argument is ", average)
	}
	io.WriteString(os.Stdout, mainString)
}
