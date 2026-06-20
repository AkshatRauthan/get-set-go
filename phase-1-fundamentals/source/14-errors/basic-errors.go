package main

import (
	"errors"
	"fmt"
)

func division[T int | float32 | float64](a, b T) (float32, error) {
	if float32(b) == float32(0) {
		return 0.0, errors.New("Error: Division By Zero")
	} else {
		ans := float32(a) / float32(b)
		return ans, nil
	}
}

func BasicErrors() {
	println("\n01. Basic Errors: ")

	res1, err1 := division(5, 3)
	if err1 != nil {
		println(err1)
	} else {
		fmt.Println("Disision for 5, 3:", res1)
	}

	res2, err2 := division(5, 0)
	if err2 != nil {
		println(err2.Error())
	} else {
		println("Disision for 5, 0:", res2)
	}
}
