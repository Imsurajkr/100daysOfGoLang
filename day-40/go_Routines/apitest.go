package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var url = []string{
	"https://google.com",
	"https://www.twitter.com",
	"https://facebook.com",
}

func fetchstatus(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	// for _, urls := range url {
	wg.Add(1)
	func() {
		resp, err := http.Get("http://google.com")
		if err != nil {
			fmt.Fprint(w, "%+v", err)
		}
		fmt.Fprint(w, "%+v", resp)
		wg.Done()
	}()
	//}
	wg.Wait()
}

func main() {
	fmt.Println("execution started ")
	http.HandleFunc("/", fetchstatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("execution ended")
}
