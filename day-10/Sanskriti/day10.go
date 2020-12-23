/*
  Command Line Arguments
  We have a package called as os package that contains an array called as “Args”.
  Args is an array of string that contains all the command line arguments passed.
*/
package main

import (
	"fmt"
	"os"
)

func main() {

	// The first argument is always program name / filepath
	myProgramName := os.Args[0]

	// this will take 4 command line arguments
	cmdArgs := os.Args[4]

	// getting the arguments with normal indexing
	gettingArgs := os.Args[2]

	toGetAllArgs := os.Args[1:]

	// it will display the program name / filepath
	fmt.Println(myProgramName)

	fmt.Println(cmdArgs)
	fmt.Println(gettingArgs)
	fmt.Println(toGetAllArgs)
}
