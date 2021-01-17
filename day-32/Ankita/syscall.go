/*
================Syscall Package==============
# this package provides a large amount of functions and types related to low level operating primitives.
#the syscall package is extensively used by other Go packages, such as os, net, and time, which all provide a portable
interface to the operating system.This means that the syscall package is not the most
portable package in the Go library – that is not its job.

==============Why syscall package?==============
it deals with incompatibilities of system internals as gently as possible and is also well documented.

==================system call======================
a system call is a programmatic way for an application to request something from the kernel of an operating system.
system calls are responsible for accessing and working with most UNIX low-level elements such as
processes, storage devices, printing data, network interfaces, and all kinds of files.
*/
package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	pid, _, _ := syscall.Syscall(39, 0, 0, 0)
	fmt.Println("My pid is", pid)
	uid, _, _ := syscall.Syscall(24, 0, 0, 0)
	fmt.Println("User ID:", uid)
	message := []byte{'H', 'e', 'l', 'l', 'o', '!', '\n'}
	fd := 1
	syscall.Write(fd, message)
	fmt.Println("Using syscall.Exec()")
	command := "/bin/ls"
	env := os.Environ()
	syscall.Exec(command, []string{"ls", "-a", "-x"}, env)
}
