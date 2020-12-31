package main

import (
	"fmt"
	"os"
)

type MyError struct{}

func (m *MyError) Error() string {
	return "boom"
}
func sayHello() (string, error) {
	return "", &MyError{}
}
func main() {
	s, err := sayHello()
	if err != nil {
		fmt.Println("unexpected error: err:", err)
		os.Exit(1)
	}
	fmt.Println("The string:", s)
}
