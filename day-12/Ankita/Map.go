package main

import (
	"fmt"
	"runtime"
)

func main() {
	var N = 400000
	var data = make(map[int]int)
	//var data = make(map[int]*int)
	for i := 0; i < N; i++ {
		value := i
		//data[value] = &value
		data[value] = value
	}
	fmt.Print(data)
	runtime.GC()
	_ = data[0]
}
