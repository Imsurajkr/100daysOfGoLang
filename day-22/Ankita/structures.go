/*
==========two structures with the same fields will not be consideredidentical in Go if their
fields are not in exactly the same order.
*/
package main

import "fmt"

type s1 struct {
	name           string
	age            int
	profile        string
	qualifications []string
}

func main() {

	//var p1 = s1{name: "Ankita", age: 22, profile: "Reactjs Developer", qualifications: ["MCA","BSC"]}
	p1 := s1{name: "Ankita", age: 22, profile: "Reactjs Developer", qualifications: []string{"Bsc", "MCA"}}
	fmt.Println(p1.qualifications)
	fmt.Println(p1)
	p2 := [9]s1{}
	p2[0] = s1{name: "harsh", age: 20, profile: "machine learning", qualifications: []string{"btech", "mtech"}}
	p2[1] = s1{name: "somi", age: 20, profile: "teacher", qualifications: []string{"bsc", "msc"}}
	fmt.Println(p2[0].qualifications)
	fmt.Println(p2[0])
	fmt.Println(p2[1].qualifications)
	fmt.Println(p2[1])
	fmt.Println(p2)
}
