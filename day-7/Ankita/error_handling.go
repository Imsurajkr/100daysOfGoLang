/*
go likes error very much and treats them as their gf that's why they have created a seperate
data type "error"
you can create yur own error message if you found that go error messages are not adequate.
having an error is one thing but how you will react to that error is a different thing.
some error needs program termination and some needs just a warning and etc.

errors in go can be returned like an object from a function or methods.

==============what is error handling============
error handling and writting error to output both are different thing
error handling is writting standard error to file descriptor
*/

package main

import (
	"errors"
	"fmt"
)

func returnError(a, b int) error {
	if a == b {
		//here we are converting an error into a string variable using errors.Error() function.
		//it will compare errror variable with a string varibale
		err := errors.New("errors in returnError() function!!!!!!!!!!")
		return err
	} else {
		return nil
	}
}

func main() {
	err := returnError(1, 2)
	if err == nil {
		fmt.Println("returnError ended normallyyyyyyyyyyyyyyyyyyyy")
	} else {
		fmt.Println(err)
	}
	err = returnError(10, 10)
	if err == nil {
		fmt.Println("returnError ended normally")
	} else {
		fmt.Println(err)
	}
}
