package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func wait_goRoutine(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)
	go func() {
		defer close(temp)
		time.Sleep(time.Second * 5)
		w.Wait()
	}()
	select {
	case <-temp:
		return false
	case <-time.After(t):
		return true
	}
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		return
	}
	var w sync.WaitGroup
	w.Add(1)
	t, err := strconv.Atoi(arguments[1])
	if err != nil {
		return
	}
	duration := time.Duration(int32(t)) * time.Millisecond
	fmt.Printf("Timeout period is %s\n", duration)
	if wait_goRoutine(&w, duration) {
		fmt.Println("Timeout")
	} else {
		fmt.Println("okkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk")
	}
	w.Done()
	if wait_goRoutine(&w, duration) {
		fmt.Println("Timeout")
	} else {
		fmt.Println("okkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk")
	}
}
