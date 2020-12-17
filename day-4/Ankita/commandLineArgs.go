package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	mainString := ""
	arguments := os.Args
	// underscore is used to ignore standard error messages
	//strnconv is used to convert string argument to number
	max, _ := strconv.ParseInt(arguments[1], 10, 64)
	min, _ := strconv.ParseInt(arguments[1], 10, 64)
	if len(arguments) == 1 {
		mainString = "Please provide one input at least"
	} else {
		for i := 0; i < len(arguments); i++ {
			n, _ := strconv.ParseInt(arguments[i], 10, 64)
			if max < n {
				max = n
			}
			if min > n {
				min = n
			}
		}
		fmt.Println("Max ", max)
		fmt.Println("Min ", min)
	}
	io.WriteString(os.Stdout, mainString)
}
