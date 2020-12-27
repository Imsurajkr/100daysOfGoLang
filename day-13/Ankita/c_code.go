/*
Go is used for a better programming experience but c is still fast and very much capable programming language.package Ankita

======Why you need to call c progrma from your go program file?================================
there are many situation while working with databse and device drivers written in c that requires your c code.
Go provides a pseudo package called 'C' to interface with c libraries.
using cgo is not always best the course to take. You’ll be losing many of the
type safety benefits and cross-compilation targets that go provides.
 It should be used as a last resort: when you’re absolutely
sure that it is not possible to achieve your objective in native go code.
*/

package main

// #include <stdio.h>
// void callC(){
// 	printf("Calling c code from go program");
// }
import "C"
import "fmt"

func main() {
	fmt.Println("go  program Statement")
	//go compiler now what to do with c code written above in commented lines
	C.callC()
	fmt.Println("Another Go statement")
}
