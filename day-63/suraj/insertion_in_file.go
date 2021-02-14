package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("Write Line one")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
