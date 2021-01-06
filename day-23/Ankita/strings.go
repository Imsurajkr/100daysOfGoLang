/*
1.strings in Go are value types
2.Go supports UTF-8 strings bydefault, which means that you do not need to load any
special packages or do anything tricky in order to print Unicode characters.
3.A Go string is a read-only byte slice that can hold any type of bytes and can have an
arbitrary length.
*/
package main

import "fmt"

func main() {
	var stingLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	fmt.Println(stingLiteral)
	fmt.Printf("x : %x\n", stingLiteral)
	fmt.Printf("sLiteral length: %d\n", len(stingLiteral))
	for i := 0; i < len(stingLiteral); i++ {
		fmt.Printf("%x", stingLiteral[i])
	}
	fmt.Println()
	fmt.Printf("%q\n", stingLiteral)
	fmt.Printf("%+q\n", stingLiteral)
	fmt.Printf("%x\n", stingLiteral)
	s2 := "€£³"
	for x, y := range s2 {
		fmt.Printf("%#U starts at byte position %d\n", y, x)
	}
	fmt.Printf("s2 length: %d\n", len(s2))
	const s3 = "ab12AB"
	fmt.Println("s3:", s3)
	fmt.Printf("x: % x\n", s3)
	fmt.Printf("s3 length: %d\n", len(s3))
	for i := 0; i < len(s3); i++ {
		fmt.Printf("%x ", s3[i])
	}
	fmt.Println()
}
