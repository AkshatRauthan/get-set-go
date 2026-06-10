package main

import "fmt"

func doSomethingSomething(m map[int]int) {
	m[4] = 4
	m[5] = 5
}

func referencedMaps() {
	m := map[int]int{1: 1, 2: 2, 3: 3}

	// For doing this simply use non-formatted print/println functions
	println("This is what a map stores internally:  ", m)
	println("Here we get the map's virtual address")
	println("Now when we pass a map to a function it works as pass by value, so the function recieves a copy of the address of the map")
	println("Which is now also the memory location the map of the function called, thus making all changes persistent")
	println("Therefore, all additions and updations will be persistent")

	fmt.Println("\nInitial map: ", m)
	doSomethingSomething(m)
	fmt.Println("Map after function call: ", m)
}
