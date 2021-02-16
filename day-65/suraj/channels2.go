package main

import (
	"fmt"
	"time"
)

func display(ch chan int) {
	time.Sleep(5 * time.Second)
	fmt.Println("Inside display()")
	ch <- 1234
}

func main() {
	ch := make(chan int)
	go display(ch)
	x := <-ch
	fmt.Println("Inside main()")
	fmt.Println("Printing x in main() after taking from channel:", x)
}
