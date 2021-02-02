package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	a, b := swap("Suraj", "Kumar")
	fmt.Println(a, b)
}
