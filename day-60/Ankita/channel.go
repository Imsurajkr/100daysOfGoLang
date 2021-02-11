package main

import (
	"fmt"
	"time"
)

func data1(ch chan string) {
	time.Sleep(4 * time.Second)
	ch <- "from data1()"
}

func data2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "from data2()"
}

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)
	go data1(chan1)
	go data2(chan2)
	select {
	case x := <-chan1:
		fmt.Println(x)
	case y := <-chan2:
		fmt.Println(y)
	}
}
