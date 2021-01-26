package main

import "fmt"

func number(num ...int) <-chan int {
	n := make(chan int)
	go func() {
		for _, i := range num {
			n <- i
		}
		close(n)
	}()

	return n
}

func square(in <-chan int) <-chan int {
	sq := make(chan int)
	go func() {
		for j := range in {
			sq <- j * j
		}
		close(sq)
	}()

	return sq
}

func main() {
	c := square(number(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println(<-c)
}
