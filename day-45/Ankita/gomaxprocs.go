package main

import (
	"fmt"
	"runtime"
)

func GOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

func main() {
	fmt.Printf("GOMAXPROCS running in this process are %d", GOMAXPROCS())
}
