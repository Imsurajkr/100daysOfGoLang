/*
Write a Go program that keeps reading integers until it gets the word END as input.
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
		if scanner.Text() == "END" {
			f.Close()
		} else {
			fmt.Println(">", scanner.Text())
		}
	}
}
