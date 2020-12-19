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
	var sum int64 = 0
	if len(os.Args) == 1 {
		mainString = "Please provide one input at least"
	} else {
		for i := 1; i < len(os.Args); i++ {
			n, _ := strconv.ParseInt(os.Args[i], 10, 32)
			sum = sum + n
		}
		fmt.Println("Sum of command line argument is ", sum)
	}
	io.WriteString(os.Stdout, mainString)
}
