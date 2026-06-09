package main

import "fmt"

// Closure functions are able to maintain their scope variables values btw calls....
// Here we have implemented a counter using closure...

// counter initialises count as 0 when called...
// It then in turn returns two internal func. one to fetch count and one to increment counts value...
// Here counter is called only once... But internal func. are maintaining the expected behavoiur

func counter() (func() int, func() int) {
	var count int = 0

	incrementCount := func() int {
		count++
		return count
	}

	getCount := func() int {
		return count
	}
	return getCount, incrementCount
}

func closure() {
	getCount, incrementCount := counter()

	fmt.Println("\nInitial Count: ", getCount())             // -> 0
	incrementCount()                                         // -> 1
	incrementCount()                                         // -> 2
	incrementCount()                                         // -> 3
	fmt.Println("Incrementing 4th time: ", incrementCount()) // -> 4
	fmt.Println("Final Counter Value: ", getCount())         // -> 4
}
