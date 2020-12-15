//package main tells the compiler that it is a executable program not a shared library.
//it is the entry point of our executable program.

package main

//import statement is used to importing a package another package.
//import fmt is imported to use function Println which comes from go standard library.
//compiler will look into GOROOT and GOPATH for imported packages.
//packages from standard library are installed in GOR OOT location while packages
//created by yourself and third party packages whichyou have installed are available in GOPATH location
import "fmt"

//implementation of code takes place here in main function.
//you can not have multiple main function in your in the files of a single project.
func main() {
	fmt.Println("Hello World!!")
}

/*Steps to install go in windows----------------------------------------------------------------
1.go to official website of go  golang.org
2.click on the documents option then many options will be listed below
3.click o installing Go option
4.very first option Go download provides a button to download the go installer.
click on that and download will be started.
5.install it by clicking next next option
6.check whether go is installed or not by typing go version in command prompt.
it will list you the version of go installed on your system.
7.if it shows you the error that go is not recognised as internal
 or external command then you should have to set path for go in users environment variable.
8.set it like --------------------
   GOPATH(variable name)     c:\Users\ankita\go
9.open vs code and write your first go program and save your file with .go extension.
10.run your first Go program in terminal by---------
	  go run HelloGo.go
  and output will be displayed like --------------------
        Hello World!!
*/
