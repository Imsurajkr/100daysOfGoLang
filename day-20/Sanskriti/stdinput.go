// Getting user input
/*
   There are 3 ways of getting user input:
   1. Reading command line
   2. By standard input
   3. By reading external files
*/

// Reading from standard input

// Here we use 'bufio' package which is mostly used for file input and output.
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

/*
   First, there is a call to bufio.NewScanner() using standard input (os.Stdin) as its
   parameter. This call returns a bufio.Scanner variable, which is used with the Scan()
   function for reading from os.Stdin line by line. Each line that is read is printed on the
   screen before getting the next one. Please note that each line that is printed by the program
   begins with the > character.
*/

// Press Control + D on a new line to end this program!
