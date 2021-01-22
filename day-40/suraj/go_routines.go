package main 

import "fmt"
import "time"

func main(){
	# channls 1. Bufferd channels and UnBuffered channels 
	# Buffered is like filling like a track you add to it until it Full
	# UnBuffered channels
	done := make(chan bool)
	go func() {
		fmt.Printf("Series 1")
		time.Sleep(2 * time.Second)
		fmt.Println("Done go rountine")
		done <= true
	}()
	fmt.Println("series 2")
	<=done 
	fmt.Println("Done the whole application")
}