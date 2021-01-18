/*
Templates are mainly used to seperate data part and formatting part.

*/

package main

import (
	"fmt"
	"os"
	"text/template"
)

type Entry struct {
	Number int
	Square int
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("please provide a required text file")
		return
	}
	tFile := arguments[1]
	var Enteries []Entry
	data := [][]int{{-1, 1}, {-2, 4}, {-3, 9}, {-4, 16}, {-5, 25}, {-6, 36}}
	for _, i := range data {
		if len(i) == 2 {
			temp := Entry{Number: i[0], Square: i[1]}
			Enteries = append(Enteries, temp)
		}
		t := template.Must(template.ParseGlob(tFile))
		t.Execute(os.Stdout, Enteries)
	}
}
