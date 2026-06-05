package main

import "fmt"

// Slices: Dynamic Sized Arrays
func slices() {
	// uninitialized slices are nil by default
	var v []int
	v = append(v, 10)
	v = append(v, 20)
	v = append(v, 30)
	v = append(v, 40)
	v = append(v, 50)

	fmt.Println("\nSlice: ", v)

	// len => length: no of elements currently in slice
	fmt.Println("Size Of Slice: ", len(v))

	// cap => capacity: current size of underlying array => doubles every time size overflows
	fmt.Println("Capacity Of Slice: ", cap(v))


	// defining non nil slices...
	var a = make([]int, 10, 12) // slice type, len, cap
	fmt.Println("\nNon nil slice: ", a)
	fmt.Println("Size Of Slice: ", len(a))
	fmt.Println("Capacity Of Slice: ", cap(a))
}
