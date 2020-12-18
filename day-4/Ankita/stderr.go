/*
   Standard error in Go

   #standard error in Go is included in stderr.go
   #writing to standard error requires use of fle descriptors related to standard errors
   #
*/

package main

import (
	"io"
	"os"
)

func main() {
	mainString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		mainString = "please enter one argument!!!!"
	} else {
		mainString = os.Args[1]
	}
	/*This code is not helpful to understand difference between standard error and standard Output
	but if you are using bash and then you will find a tick to run it on your shell
		go run stderr.go 2>/temp/stdError ------ it will redirect your error
		go run stderr.go 2>/dev/null ------ it will tell your compiler to completely ignores the error
	To merge standard output and standard error you can use following command line
	go run stderr.go >/temp/output 2>&1 -------it wll merge your output stream and error stream
	*/
	io.WriteString(os.Stdout, "first day with Stderr\n")
	io.WriteString(os.Stderr, mainString)
}
