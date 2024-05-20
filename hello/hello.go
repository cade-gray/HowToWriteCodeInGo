package main // Package declaration. Every Go program starts with a package declaration. Packages are Go's way of organizing and reusing code. There are two types of Go programs: executables and libraries. Executable applications are the kinds of programs that we can run directly from the terminal. Libraries are collections of code that we package together so that we can use them in other programs. The package "main" tells the Go compiler that the package should compile as an executable program instead of a shared library. The main function in the package "main" will be the entry point of our executable program.
import (
	"example/user/hello/morestrings"
	"fmt"
) // Package implementing formatted I/O. It's like java's System.out.println and System.in Scanner
func main() {
	fmt.Println("Hello, Cade! You wrote this in Go when reading How to Code in Go. Saved in %USERPROFILE%/go/bin/hello.exe.")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
}
