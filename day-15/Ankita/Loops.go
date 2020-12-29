/*
break statement terminate program while continue terminates loop.
Go does not offer while loop but uses for loop instead of while loop.
*/
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	ank := true
	// while (ank) {
	// 	fmt.Println("entered in while loop");
	// 	ank = false
	// }
	for ank := true; ank; ank = false {
		fmt.Println("entered in while loop")
	}
}
