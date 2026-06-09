package main

import "fmt"

// Variadic functions in Go are functions that can recieve n number of arguments in them....
// Like fmt.Println etc...

// Taking n number of integers as input
func sum(v ...int) int {
	sum := 0
	for _, val := range v {
		sum += val
	}
	return sum
}

// Taking n number of strings or integers as input
func sumOrAppend(v ...any) (int, string) {
	sum := 0
	str := ""

	for _, val := range v {
		switch newVal := val.(type) {
		case int:
			sum += newVal
		case string:
			str += newVal
		}
	}

	return sum, str
}

func variadicFunctions() {
	// Using variadic functions..............

	res1 := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	res2, str := sumOrAppend(1, "A", 2, "L", 3, "A", 4, "N")

	fmt.Println("res1: ", res1)
	fmt.Println("res2: ", res2)
	fmt.Println("str: ", str)
}
