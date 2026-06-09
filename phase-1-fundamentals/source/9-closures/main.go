package main

import "fmt"

// RUN COMMAND:
// go run main.go closures.go closures-example.go

func main(){
	fmt.Println("\nClosures are functions that are able to retain values of variables present in their scope btw diffrent calls")
	closure()

	fmt.Println("\n\nExample uage of closures for greeting persons....")
	greetPersons()
}