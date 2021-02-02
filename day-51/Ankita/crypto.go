package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "my name is ankita"
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Print(s)
	fmt.Print("%x", bs)

}
