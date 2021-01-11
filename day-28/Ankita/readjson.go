package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// type Record struct {
// 	Name    string
// 	Surname string
// 	Tel     []Telephone
// }
// type Telephone struct {
// 	Mobile bool
// 	Number string
// }

func saveToJSON(filename *os.File, key interface{}) {
	fmt.Println(key, "key!!!!!!!!!!")
	fmt.Println(filename)

	encodeJSON := json.NewEncoder(filename)
	err := encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func main() {
	myRecord := Record{
		Name:    "Ankita",
		Surname: "Yadav",
		Tel: []Telephone{Telephone{Mobile: true, Number: "987654321"},
			Telephone{Mobile: true, Number: "12345670987"},
			Telephone{Mobile: false, Number: "a555555555"},
			Telephone{Mobile: true, Number: "567432678"},
			Telephone{Mobile: true, Number: "@@@@@@@@@@@@@@"},
		},
	}
	saveToJSON(os.Stdout, myRecord)
}
