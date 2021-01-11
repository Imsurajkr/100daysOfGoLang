package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args[1]
	var rec map[string]interface{}
	filedata, err := ioutil.ReadFile(arguments)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal([]byte(filedata), &rec)
	for key, value := range rec {
		fmt.Println("key:", key, "value:", value)
	}
}
