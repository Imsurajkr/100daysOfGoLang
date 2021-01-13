package main

import (
	"fmt"
	"os"
)

func test(s ...string) {
	for i, a := range s {
		fmt.Println(i, a)
	}
}

func main() {
	argument := os.Args
	if len(argument) == 1 {
		return
	}
	test(argument...)
	test("Ankita", "somi", "harsh", "mansi", "deepika")
}
