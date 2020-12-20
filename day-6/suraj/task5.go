package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// The *os.File is of
	var f *os.File
	f = os.Stdin
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
		if scanner.Text() == "END" || scanner.Text() == "end" {
			fmt.Println("The Program stops Now")
			f.Close()
		}
	}

}
