package main

import (
	"flag"
	"fmt"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
}

func main() {
	n := flag.Int("b", 10, "Number of goroutines")
	flag.Parse()
	count := *n
	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}
	go test()
	go func() {
		for i := 20; i < 30; i++ {
			fmt.Print(i, " ")
		}
	}()
	time.Sleep(5 * time.Second)
	fmt.Println()
}
