package main
import
	"fmt"

func main(){
	fmt.Println("hello")
var name string="Mansi Dhingra"
//variable named as name whose datatype is string is declared and passed the value "Mansi Dhingra "
//or we can simply declare any variable name with its data type and change it later on
//name="mansi"
//name="lily"
//now that will change the valuefrom mansi to lily
fmt.Println("hello",name)
fmt.Print("hello",name)

// the major difference between fmt.Println anf Printf is that fmt.Println()adds a newline character each time you call it.
fmt.Printf("%s %s","hello" , name)
//difference between fmt.Println() and fmt.Printf() is
// that the latter requires a format specifier for each thing that you want to print.


}

