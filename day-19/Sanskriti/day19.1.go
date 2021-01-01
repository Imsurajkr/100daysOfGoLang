/*
   With the help of Before() and After() and Equal(), function we can compare the time
   as well as date but we are also going to use the time.Now() and time.Now().Add() function for comparison.

   Functions Used: These functions compares the times as seconds.

   Before(temp) – This function is used to check if the given time is before
   the temporary time variable and return true if the time variable comes before
   the temporary time variable else false.

   After(temp) – This function is used to check if given time is after the temporary
   time variable and return true if time variable comes after temporary time variable else false.

   Equal(temp) – This function is used to check if given time is equal the temporary time
   variable and return true if time variable equals temporary time variable else false.
*/

package main

import (
	"fmt"
	"time"
)

// importing time module

// Main function
func main() {

	today := time.Now()
	tomorrow := today.Add(24 * time.Hour)

	// Using time.Before() method
	g1 := today.Before(tomorrow)
	fmt.Println("today before tomorrow:", g1)

	// Using time.After() method
	g2 := tomorrow.After(today)
	fmt.Println("tomorrow after today:", g2)

}
