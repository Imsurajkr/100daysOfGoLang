package main

import (
	"fmt"
	"os"
)

func main(){
	i:=0
	for{
		i++
		fmt.Println("Helllooooooo")
		if i==10{
			break
		}
	}
	for i:=0;i<10;i++{
		fmt.Println("worlddddddddddddddddd")
	}
	for i:=true;i;i=false{
		fmt.Println(i)
	}
	arguments := os.Args
	if len(arguments) ==1{
       fmt.Println("no argument provided")
	}else if len(arguments) == 2{
		fmt.Println("one argument provided")
	}else{
     fmt.Println("many arguments provided")
	}
	set := true
	//no ternary support to go
	// kill := set?"hello":"Ankita"
	// fmt.Println(kill)
}