package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main function starts here")
	c := make(chan string)
	go func() {
		time.Sleep(time.Second * 5)
		c <- "it's c channel"
	}()
	select {
	case res := <-c:
		fmt.Println(res)
	case <-time.After(6 * time.Second):
		fmt.Println(time.Second)
	}
}
