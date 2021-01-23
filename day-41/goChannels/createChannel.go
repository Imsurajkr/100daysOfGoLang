package main

import (
	"fmt"
	"time"
)

func writeChannel(c chan<- int, x int) {
	fmt.Println(x, "@@@@@@@@@@@@@@@@@@@2")
	c <- x
	c <- 100
	c <- -60
	close(c)
	fmt.Println(x, "AAAAAAAAAAAAAAAAAAA")
}

func main() {
	c := make(chan int)
	go writeChannel(c, 10)
	time.Sleep(5 * time.Second)
	fmt.Println("Print ", <-c)
	fmt.Println("Print ", <-c)
	fmt.Println("Print ", <-c)
	time.Sleep(2 * time.Second)
	read, ok := <-c
	fmt.Println(read)
	if ok {
		fmt.Println("channel open")
	} else {
		fmt.Println("Channel closed")
	}
}
