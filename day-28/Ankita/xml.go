package main

import (
	"fmt"
	"os"
)

func main() {

	// Open our xmlFile
	xmlFile, err := os.Open("users.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

}
