package main

import "fmt"

func main() {
	var mapLit map[string]int // making map
	var mapAssigned map[string]int
	mapLit = map[string]int{"one": 1, "two": 2} // adding key-value pair
	mapCreated := make(map[string]float32)      // making map with make()
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5 // creating key-value pair for map
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3 // changing value of already existing key
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
}
