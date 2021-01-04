// Reading from command Line
package main

/* Here 'os' package is used for getting command-line argument.
   'Strconv' package is used to convert command-line argumnet which is
   basically string into arithmetical data type.

*/
import (
	"fmt"
	"os"
	"strconv"
)

// Here we check whether the user has given any input or not
// If user didn't give any input the below statement is printed on the terminal.
func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}
	// Here os.Args is a Go slice with string values.
	// Where the first element is the excecuteable program.
	// So in order to find min and max float value, we need
	// second element of the string os.Args which is index[1].

	arguments := os.Args

	// We use strconv.ParseFloat() function to ignore the error
	// By the given below statement we tell Go to use only first value
	// return by strconv.ParseFloat() function and don't want the second value
	// which is assign to '_' , which is a blnak identifier (which is the Go way of discarding a value).
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	// And we use for loop to iterate all the elements of os.Args slice.
	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
