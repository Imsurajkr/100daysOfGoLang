package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	go func(c chan string) {
		time.Sleep(2 * time.Second)
		c <- "result 1"
	}(c1)
	select {
	case data := <-c1:
		fmt.Println(data)
	case <-time.After(1 * time.Second):
		fmt.Println("timed out")
	}
	c2 := make(chan string)
	go func(c chan string) {
		time.Sleep(2 * time.Second)
		c <- "result 1"
	}(c2)
	select {
	case data := <-c1:
		fmt.Println(data)
	case <-time.After(3 * time.Second):
		fmt.Println("timed out")
	}
}
