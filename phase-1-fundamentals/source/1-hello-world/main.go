package main

// declaring the package name as "main".
// The package name is used to organize and reuse code in Go.
// It allows us to group related functions, types, and variables together.
// In this case, we are declaring a package named "main" which can be imported and used in other Go files.

import "fmt"

func main() {
	fmt.Println("hello world") // Adds a new line at the end of the output
	fmt.Print("hello world")   // Adds spaces btw non-strings arguments
	fmt.Printf("hello world")  // For using format specifiers (%d, %s etc) to format the output
}
