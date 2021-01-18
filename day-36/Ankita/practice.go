package main
import (
	"fmt"
	"time"
)
const val = "Yadav"
const num float64 = 800000 
// func main(){
// 	val := "ANkita"
// 	fmt.Printf("%T",num)
// 	fmt.Println(float64(num))
// 	fmt.Println(6e267/num)
// 	fmt.Println(math.Sqrt(num))
// 	fmt.Println(val)
// 	fmt.Println("hello world")
// 	fmt.Println("Ankita"+" Yadav")
// 	fmt.Println("1+1= ",1+1)
// 	fmt.Println("2.8888/2= ",2.8888/2)
// 	fmt.Println(true&&false)
// 	fmt.Println(true||false)
// 	fmt.Println(!true)
// 	//not defined on bools
// 	// fmt.Println(true&false)
// 	// fmt.Println(true|false)
// 	var e uint;
// 	fmt.Println(e)
// 	nme := "Ankita"
// 	fmt.Println(nme)
// 	var a,b int = 8,654
// 	fmt.Println(a,b)
// 	boolean := "true"
// 	fmt.Printf("%T",boolean)
// }

func main(){
	 i:=5;
	switch i{
	case 1:
		fmt.Println("one case")
	case 3:
		fmt.Println("third case")
	case 5:
		fmt.Println("fifth case")
	case 7:
		fmt.Println("seven cases")
	default:
		fmt.Println("default")
	}
	fmt.Println("")
	switch time.Now().Weekday() {
		case time.Sunday:
			fmt.Println("Sunday")
		case time.Monday:
			fmt.Println("Monday")
		case time.Tuesday:
			fmt.Println("Tuesday")
		case time.Wednesday:
				fmt.Println("Wednesday")
		case time.Thursday:
					fmt.Println("Thursday")
		case time.Friday:
					fmt.Println("FRiday")
		case time.Saturday:
					fmt.Println("Saturday")
					
	}
	fmt.Println("")
	t:=time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() > 12:
		fmt.Println("Good Evening")
	}
	cond := func(i interface{}){
		switch i.(type){
		case bool:
			fmt.Println("boolean")
		case int:
			fmt.Println("integer")
		
	    case string:
		  fmt.Println("string")
	}
	}
	cond(15)
	cond("ANkita")
}