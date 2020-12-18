
// The Infinite Loop
// A loop becomes an infinite loop if its condition never becomes false.

package main

import "fmt"

func main() {
//    for true  {
//        fmt.Printf("This loop will run forever.\n");
//    }
   /* local variable definition */
   var a int = 100
   var b int = 200
   var ret int

   /* calling a function to get max value */
   ret = max(a, b)

   fmt.Printf( "Max value is : %d\n", ret )
}

/* function returning the max between two numbers */
func max(num1, num2 int) int {
   /* local variable declaration */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result 
}
