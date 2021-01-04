package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func k1() {
	myVersion := runtime.Version()
	fmt.Println("----- That's My Version ----", myVersion)
	major := strings.Split(myVersion, ".")[0][2]
	fmt.Println("----- That's Major ----", major)
	minor := strings.Split(myVersion, ".")[1]
	fmt.Println("----- That's Minor ----", minor)
	m1, _ := strconv.Atoi(string(major))
	fmt.Println("----- let's check m1 ----", m1)
	m2, _ := strconv.Atoi(minor)
	fmt.Println("----- let's check m2 ----", m2)
	if m1 == 1 && m2 < 8 {
		fmt.Println("Need Go version 1.8 or higher!")
		return
	}
	fmt.Println("You are using Go version 1.8 or higher!")
}
