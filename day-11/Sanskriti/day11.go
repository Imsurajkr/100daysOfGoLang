// Main() function
/*
	The main() function is a special type of function and it is the entry point of the executable programs.
	It does not take any argument nor return anything.
	Go automatically call main() function, so there is no need to call main() function explicitly
	and every executable program must contain single main package and main() function.
*/

// Program using main() fuction

// Declaration of the main package
package main

// Importing packages
import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// Main function
func main() {

	// Sorting the given slice
	s := []int{345, 78, 123, 10, 76, 2, 567, 5}
	sort.Ints(s)
	fmt.Println("Sorted slice: ", s)

	// Finding the index
	fmt.Println("Index value: ", strings.Index("Hello_There!How_you_doin?", "lo"))

	// Finding the time
	fmt.Println("Time: ", time.Now().Unix())

}
