package main

import "fmt"

func main() {
	var a string = "string"
	b := "string"
	//both the variables are of string type ! , and varible can be assing in any of the formate
	if a == b {
		fmt.Println("same")
	}

	// b = 56
	// this will give us error as we cant type cast the variables into other datadype
	// neither we use ":" before "=" ,i.e b:=56 as b is alredy decleared and used .
}
