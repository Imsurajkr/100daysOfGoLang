/*
UNIX consider everything as file whether it is your mouse,a file ,a pointer or anything.
UNIX uses file decriptors.
what are file descriptors?
they are just positive vaues that uniquely identifies an open filein computer os.
UNIX uses three special and standard file names
   /dev/stdin          standard input           descriptor-0
   /dev/stdout         standard oputput          descriptor-1
	/dev/stderr        standard error            descriptor-2

	go uses os.Stderr,os,Stdin,os.Stdout for accessing standard input,output and error
	you can use  /dev/stderr  , /dev/stdin and  /dev/stdout  but it is better,safer and more portable to
	stick with os.Stderr standards

	Printing Output----------------------------------------------------------------

	all the below mentioned functions are present in "fmt" package.
	printf()--- use when you want to modify output according to you like if you want to add modifier specifiers.
	println()--use when you want to add a new line character at the end of line automatically
	print()--you can use When want to print anything without any format specifier and new line character

	S family of print functions-------create strings based on a given format
	Sprintf()
	Sprintln()
	Sprint()

	F family of print function--------used for writing to files using an io.writer
	Fprint()
	Fprintf()
	Fprintln()
Using Standard Output

standard output requires that do  not belong to fmt package
*/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
	var age = 22
	var name = "Ankita"

	fmt.Println(name)
	fmt.Println(age)
	fmt.Print(name)
	fmt.Print(age)
	fmt.Printf("%d", age)
	fmt.Printf("%s", name)
}
