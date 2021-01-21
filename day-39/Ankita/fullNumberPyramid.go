package main

import (
	"fmt"
)

func spaces(i int) {
	if i == 0 {
		return
	}
	spaces(i - 1)
	fmt.Print("  ")
}

func num(i int, k int, cond bool) {
	if i == 0 || i == k {
		return
	} else if cond == true {
		num(i-1, k, cond)
		fmt.Print(i, " ")
	} else {
		fmt.Print(i, " ")
		num(i-1, k, cond)
	}

}

func main() {
	j := 1
	k := 4
	for i := 0; i < 5; i++ {
		cond := true
		spaces(k)
		num(i+j, i, cond)
		cond = false
		num(i+j-1, i, cond)
		spaces(4)
		j = j + 1
		k = k - 1
		fmt.Println("")
	}
}
