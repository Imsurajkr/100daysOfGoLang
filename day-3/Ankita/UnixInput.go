/*
GETTING USER INPUT

:= --------------------------------
it is defined as short-hand assignment operator
it is used to assign values to variables without using var keyword .
it will not reassign values to already defined var and if you do so it will throw an error.
it can not be used outside a function but can be used inside a function.

= --------------------------------
it is used to assign values to variables using var keyword
it will reassign values to already defined var.package Ankita
it can be used outside function.

STANDARD USER INPUT --------------------------------

buffio package is mostly used for file input and output
and its common functionality is to access command line arguments of Go program.

OS package------------------------------------------------
1.it provides functions for creating,deleting,renamng files and directories as well as functions for learning the
UNIX permissions and other directories.
2.its funtion will work on both Unix and Windows.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}
