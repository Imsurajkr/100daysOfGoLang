package main
import (
	"fmt"
	"math"
)
const val = "Yadav"
const num float64 = 800000 
func main(){
	val := "ANkita"
	fmt.Printf("%T",num)
	fmt.Println(float64(num))
	fmt.Println(6e267/num)
	fmt.Println(math.Sqrt(num))
	fmt.Println(val)
	fmt.Println("hello world")
	fmt.Println("Ankita"+" Yadav")
	fmt.Println("1+1= ",1+1)
	fmt.Println("2.8888/2= ",2.8888/2)
	fmt.Println(true&&false)
	fmt.Println(true||false)
	fmt.Println(!true)
	//not defined on bools
	// fmt.Println(true&false)
	// fmt.Println(true|false)
	var e uint;
	fmt.Println(e)
	nme := "Ankita"
	fmt.Println(nme)
	var a,b int = 8,654
	fmt.Println(a,b)
	boolean := "true"
	fmt.Printf("%T",boolean)
}