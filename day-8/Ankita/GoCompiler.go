/*
Go compiler is executed with go tool
will go compiler only generate executable file?
No,they do many more things than them like they generate archive file,object file,race conditions,etc.
##### go tool compile <program_name>.go
it will produce object code (machine code in relocatable format)
???????advantage of using relocatable format------requires low memoryas possible during the linking phase.
#########go tool compile -pack <program_name>.go
it will produce archive file(binary file that contains one or more than one file)
You can list the contents of an .a archive file as follows:
######## ar t unsafe.a
__.PKGDEF
_go_.o
*/

package main

import "fmt"

func main() {
	fmt.Println("Hello World!!")
}
