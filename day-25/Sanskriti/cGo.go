package main

/*
   We include C code in comments of the program, because the Go tool knows
   how to handle such comments because we are using "C Go Package"
*/

//#include <stdio.h>
//void callC() {
// printf("Calling C code!\n");
//}
import "C"

// And the following code of program is Go code, so we import the other package seperately.

import "fmt"

func main() {
	fmt.Println("A Go statement!")
	C.callC()
	fmt.Println("Another Go statement!")
}

//In order to execute the callC() C function, you will need to call it as C.callC().
