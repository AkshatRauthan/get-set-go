package main

import "fmt"

// zeroed values
// int -> 0, bool -> false, string -> "", pointers -> nil

func arrays1D() {
	// Arrays in Go are fixed size, they cannot be resized after declaration
	var v [5]int
	for i := range len(v) {
		v[i] = i
	}
	fmt.Println("Length Of Array: ", len(v))
	fmt.Println("Values Of Array: ", v)

	// Declaring Arrays with Initialization
	a := [5]string{"Go", "Is", "A", "Great", "Language"}
	fmt.Println("\nString Array: ", a)

	// Iterating over an array
	fmt.Println("\nIterating Over An Array: ")
	for i := range v {
		fmt.Printf("Index: %d, Value: %d\n", i, v[i])
	}
}
