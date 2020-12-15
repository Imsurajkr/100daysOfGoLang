/*
  Compiling Go code

  write your go code save it with .go extension and
	to  create a statically linked executable file,you will need to execute the following command:
	go build <fileName>.go
	after that you have a new executable file that you will need to execute:
	file <file_Name>
		<file_Name>: Mach-O 64-bit executable x86_64
		$ ls -l <file_Name>
		-rwxr-xr-x 1 mtsouk staff 2007576 Dec 14 21:10 <file_Name>
		$ ./<file_Name>
		This is a sample Go program!


The main reason why the file size of <file_Name>is that big is because it is staticallylinked,
 which means that it does  not require any external libraries to run.


		 Executing Go code

		 if you want to run your code without creating any permanent executable file
		 then run following command
		  go run <source_file_name>.go
		But do you wonder how go compiler will run code without any error without creating any
		executable files ?
		the answer is that compiler will create an executable file but it automatically execute and
		automatically deletes after the program has finished execution.package Ankita
		so,you might think that no executable file is created by the compiler.
		one major advantage of go run over go build is that
		go run never leaves any file on hard disk after te program has finished its execution.



		TWO GO RULES
		1.you can either use a go package or you do not include it

			 if you include any package that you won't use in future will throw you error.
			 you can include those packages with a prefix underscore sign and they will not throw any error.

		2.There is only one way to format curly braces



		Downloading Go packages

		You can not simply import a package and run your program ,it will show you an error message.

		you have to download that package on your computer using following command
		go get -v <package_name>
		it will show you-------------
			github.com/mactsouk/go (download)
			github.com/mactsouk/go/simpleGitHub
			#to check details about downloaded package run following command
			 ls -l ~/go/src/github.com/mactsouk/go/simpleGitHub
			 #go get command also compiles the package ,so if you want to check the compiled package then run following command
			  ls -l ~/go/pkg/windows_amd64/github.com/mactsouk/go/simpleGitHub.a
			#to delete package  run following command
				 go clean -i -v -x github.com/mactsouk/go/simpleGitHub
				 it shows
				 cd C:\Users\ankita\go\src\github.com\mactsouk\go\simpleGitHub
					rm -f simpleGitHub.test simpleGitHub.test.exe
					rm -f C:\Users\ankita\go\pkg\windows_amd64\github.com\mactsouk\go\simpleGitHub.a
					and deletes the package
*/

package main

import (
	"fmt"

	"github.com/mactsouk/go/simpleGitHub"

	//including package that is not used throws error
	// "os"
	//here it will not throw any error
	_ "os"
)

func main() {
	fmt.Println("welcome to day-2 learning")
	//it will add two numbers
	fmt.Println(simpleGitHub.AddTwo(4, 5))
}
