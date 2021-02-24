package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsAny("Germany", "G"))
	fmt.Println(strings.ContainsAny("Germany", "g"))

	fmt.Println(strings.Contains("Germany", "Ger"))
	fmt.Println(strings.Contains("Germany", "ger"))
	fmt.Println(strings.Contains("Germany", "er"))

	fmt.Println(strings.Count("cheese", "e"))

	fmt.Println(strings.EqualFold("Cat", "cAt"))
	fmt.Println(strings.EqualFold("India", "Indiana"))
}
