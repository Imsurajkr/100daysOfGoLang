package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Not enough arguments!")
		return
	}
	for _, file := range os.Args[1:] {
		fmt.Println("Processing:", file)
		f, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		One := token.NewFileSet()
		files := One.AddFile(file, One.Base(), len(f))
		var myScanner scanner.Scanner
		myScanner.Init(files, f, nil, scanner.ScanComments)
		for {
			pos, tok, lit := myScanner.Scan()
			if tok == token.EOF {
				break
			}
			fmt.Printf("%s\t%s\t%q\n", One.Position(pos), tok, lit)
		}
	}
}
