package main

import "fmt"

func main() {
	err := Run()
	if err != nil {
		fmt.Printf("Error: %#v\n", err)
	}
}
