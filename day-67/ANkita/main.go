package main

import (
	b "ANkita/basic"
	"ANkita/basic/gross"

	"fmt"
)

func main() {
	b.Basic = 10000
	fmt.Println(gross.GrossSalary())
}
