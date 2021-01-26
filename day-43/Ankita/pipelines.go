/*
Create a pipeline that reads text files, finds the number of occurrences of a given
phrase in each text file, and calculate the total number of occurrences of the
phrase in all files.
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var count map[string]int = make(map[string]int)

func fileRead(fileName *os.File) []string {
	var l []string
	r := bufio.NewReader(fileName)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			fmt.Println(err)
			break
		} else if err != nil {
			fmt.Println("Error opening file", err)
		}
		fmt.Println(line)
		for _, line := range strings.Fields(line) {
			l = append(l, line)
		}
	}
	return l
}

func phraseCount(filename []string, phrase string) map[string]int {
	count[phrase] = 0
	for _, word := range filename {
		_, ok := count[word]
		if ok {
			count[phrase]++
		} else {

			continue
		}
	}
	return count
}

func main() {
	argument := os.Args
	if len(argument) == 1 {
		fmt.Println("you have to provide a file name to find a phrase")
		return
	} else if len(argument) == 2 {
		fmt.Println("you have to provide a phrase name which you want to search in given file")
		return
	}

	file, err := os.Open(argument[1])
	defer file.Close()
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	count := phraseCount(fileRead(file), argument[2])
	fmt.Println(argument[2], " existed ", count, " times in provided file ", argument[1])
}
