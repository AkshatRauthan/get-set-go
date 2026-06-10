package main

import "fmt"

func doSomethingSpecial(v []int) {
	v[0] = 67
	v[1] = 69
	v = append(v, 33)
	v = append(v, 19)
	v = append(v, 10)

	fmt.Println("Slice in called function just before reallocation: ", v)

	v = append(v, 101)
	v = append(v, 696)
	v = append(v, 969)
	v[2] = 96
	v[1] = 101

	fmt.Println("Final value of slice in called function: ", v)
}

func referencedSlices() {
	v := make([]int, 3, 6) // making a slice with len 3 and cap 6

	fmt.Println("\nInitial slice: ", v)
	fmt.Println("Capacity Of slice: ", cap(v))
	doSomethingSpecial(v)

	fmt.Println("Slice after function call: ", v)

}
