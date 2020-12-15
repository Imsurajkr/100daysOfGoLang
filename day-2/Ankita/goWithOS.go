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

*/