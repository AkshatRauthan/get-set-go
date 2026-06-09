package main

import "fmt"

// RUN COMMAND
// go run main.go functions.go first-class-citizens.go variadic-functions.go

func main() {
	functions()

	fmt.Println("\n\nFunctions in go are first class citizens..... ")
	fmt.Println("It means that they can be assigned to a variable")
	fmt.Println("Also they can be passed on to or returned from any other function just like variables")

	functionsAsFirstClassCitizens()

	fmt.Println("\n\nVariadic functions in Go are functions that can recieve any number of arguments in them")
	variadicFunctions()

}
