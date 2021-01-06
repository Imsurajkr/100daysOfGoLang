package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage:selectColumn column <file1> [<file2> [...<fileN]]")
		os.Exit(1)
	}
	temp, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("")
		os.Exit(1)
	}
	column := temp
	if column < 0 {
		fmt.Println("")
		os.Exit(1)
	}
	for _, fileName := range arguments[2:] {
		fmt.Println("\t\t", fileName)
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening file")
			continue
		}
		defer f.Close()
		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println("Error opening file", err)
			}
			l := strings.Fields(line)
			if len(l) > column {
				fmt.Println(l[column-1])
			}

		}
	}
}
