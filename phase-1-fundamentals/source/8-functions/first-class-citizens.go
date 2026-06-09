package main

import (
	"fmt"
	"math"
)

// A function returning an another function
func mask() func(int, int) int {
	return func(m, n int) int {
		return int(math.Pow(float64(m), float64(n)))
	}
}

func functionsAsFirstClassCitizens() {
	powFunc := mask()
	fmt.Println("\nUsing functions that are returned by another function:\n2^25 = ", powFunc(2, 25))
}
