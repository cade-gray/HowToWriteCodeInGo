package main // Package declaration. Every Go program starts with a package declaration. Packages are Go's way of organizing and reusing code. There are two types of Go programs: executables and libraries. Executable applications are the kinds of programs that we can run directly from the terminal. Libraries are collections of code that we package together so that we can use them in other programs. The package "main" tells the Go compiler that the package should compile as an executable program instead of a shared library. The main function in the package "main" will be the entry point of our executable program.
import (
	"example/user/hello/morestrings"
	"fmt" // Package implementing formatted I/O. It's like java's System.out.println and System.in Scanner

	// External Packages
	"github.com/google/go-cmp/cmp" // Had to run go mod tidy to import and make linter happy. It's a package that helps you compare two values and produce a human-readable diff between them. It's useful for tests and for debugging.
)

func main() {
	fmt.Println("Hello, Cade! You wrote this in Go when reading How to Code in Go. Saved in %USERPROFILE%/go/bin/hello.exe.") // I am on Windows, so it's %USERPROFILE% instead of $HOME.
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
