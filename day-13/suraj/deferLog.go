package main

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE = "/tmp/mGo.log"

func one(aLog *log.Logger) {
	aLog.Println("-- FUNCTION one ------")
	defer aLog.Println("-- FUNCTION one ------")

	for i := 0; i < 10; i++ {
		defer aLog.Println(i, "---")
		aLog.Println(i)
	}
}

func two(aLog *log.Logger) {
	aLog.Println("---- FUNCTION two")
	defer aLog.Println("FUNCTION two ------")
	for i := 10; i > 0; i-- {
		aLog.Println(i)
	}
}

// Change three to main

func three() {
	f, err := os.OpenFile(LOGFILE,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	iLog := log.New(f, "logDefer ", log.LstdFlags)
	iLog.Println("*****************")
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")
	one(iLog)
	two(iLog)
}

// cat /tmp/mGo.log
// logDefer 2020/12/26 14:28:06 Hello there!
// logDefer 2020/12/26 14:28:06 Another log entry!
// logDefer 2020/12/26 14:28:06 -- FUNCTION one ------
// logDefer 2020/12/26 14:28:06 0
// logDefer 2020/12/26 14:28:06 1
// logDefer 2020/12/26 14:28:06 2
// logDefer 2020/12/26 14:28:06 3
// logDefer 2020/12/26 14:28:06 4
// logDefer 2020/12/26 14:28:06 5
// logDefer 2020/12/26 14:28:06 6
// logDefer 2020/12/26 14:28:06 7
// logDefer 2020/12/26 14:28:06 8
// logDefer 2020/12/26 14:28:06 9
// logDefer 2020/12/26 14:28:06 -- FUNCTION one ------
// logDefer 2020/12/26 14:28:06 ---- FUNCTION two
// logDefer 2020/12/26 14:28:06 10
// logDefer 2020/12/26 14:28:06 9
// logDefer 2020/12/26 14:28:06 8
// logDefer 2020/12/26 14:28:06 7
// logDefer 2020/12/26 14:28:06 6
// logDefer 2020/12/26 14:28:06 5
// logDefer 2020/12/26 14:28:06 4
// logDefer 2020/12/26 14:28:06 3
// logDefer 2020/12/26 14:28:06 2
// logDefer 2020/12/26 14:28:06 1
// logDefer 2020/12/26 14:28:06 FUNCTION two ------
// logDefer 2020/12/26 14:31:06 -------------------
// logDefer 2020/12/26 14:31:06 Hello there!
// logDefer 2020/12/26 14:31:06 Another log entry!
// logDefer 2020/12/26 14:31:06 -- FUNCTION one ------
// logDefer 2020/12/26 14:31:06 0
// logDefer 2020/12/26 14:31:06 1
// logDefer 2020/12/26 14:31:06 2
// logDefer 2020/12/26 14:31:06 3
// logDefer 2020/12/26 14:31:06 4
// logDefer 2020/12/26 14:31:06 5
// logDefer 2020/12/26 14:31:06 6
// logDefer 2020/12/26 14:31:06 7
// logDefer 2020/12/26 14:31:06 8
// logDefer 2020/12/26 14:31:06 9
// logDefer 2020/12/26 14:31:06 -- FUNCTION one ------
// logDefer 2020/12/26 14:31:06 ---- FUNCTION two
// logDefer 2020/12/26 14:31:06 10
// logDefer 2020/12/26 14:31:06 9
// logDefer 2020/12/26 14:31:06 8
// logDefer 2020/12/26 14:31:06 7
// logDefer 2020/12/26 14:31:06 6
// logDefer 2020/12/26 14:31:06 5
// logDefer 2020/12/26 14:31:06 4
// logDefer 2020/12/26 14:31:06 3
// logDefer 2020/12/26 14:31:06 2
// logDefer 2020/12/26 14:31:06 1
// logDefer 2020/12/26 14:31:06 FUNCTION two ------
// logDefer 2020/12/26 14:33:40 *****************
// logDefer 2020/12/26 14:33:40 Hello there!
// logDefer 2020/12/26 14:33:40 Another log entry!
// logDefer 2020/12/26 14:33:40 -- FUNCTION one ------
// logDefer 2020/12/26 14:33:40 0
// logDefer 2020/12/26 14:33:40 1
// logDefer 2020/12/26 14:33:40 2
// logDefer 2020/12/26 14:33:40 3
// logDefer 2020/12/26 14:33:40 4
// logDefer 2020/12/26 14:33:40 5
// logDefer 2020/12/26 14:33:40 6
// logDefer 2020/12/26 14:33:40 7
// logDefer 2020/12/26 14:33:40 8
// logDefer 2020/12/26 14:33:40 9
// logDefer 2020/12/26 14:33:40 9 ---
// logDefer 2020/12/26 14:33:40 8 ---
// logDefer 2020/12/26 14:33:40 7 ---
// logDefer 2020/12/26 14:33:40 6 ---
// logDefer 2020/12/26 14:33:40 5 ---
// logDefer 2020/12/26 14:33:40 4 ---
// logDefer 2020/12/26 14:33:40 3 ---
// logDefer 2020/12/26 14:33:40 2 ---
// logDefer 2020/12/26 14:33:40 1 ---
// logDefer 2020/12/26 14:33:40 0 ---
// logDefer 2020/12/26 14:33:40 -- FUNCTION one ------
// logDefer 2020/12/26 14:33:40 ---- FUNCTION two
// logDefer 2020/12/26 14:33:40 10
// logDefer 2020/12/26 14:33:40 9
// logDefer 2020/12/26 14:33:40 8
// logDefer 2020/12/26 14:33:40 7
// logDefer 2020/12/26 14:33:40 6
// logDefer 2020/12/26 14:33:40 5
// logDefer 2020/12/26 14:33:40 4
// logDefer 2020/12/26 14:33:40 3
// logDefer 2020/12/26 14:33:40 2
// logDefer 2020/12/26 14:33:40 1
// logDefer 2020/12/26 14:33:40 FUNCTION two ------
