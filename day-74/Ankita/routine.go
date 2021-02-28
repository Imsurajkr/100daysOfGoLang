package main

import "fmt"

func transmitMessage(message string) {
	fmt.Println("Message Transmitted ", message)
}

func main() {
	go transmitMessage("Welcome to MilkyWay galaxy")
	fmt.Scanln()
	fmt.Println("Is Message transmitted??")
}
