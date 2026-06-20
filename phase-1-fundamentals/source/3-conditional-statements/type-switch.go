package main

import "fmt"

func TypeSwitch() {
	// Using Type switch
	mul := func(i any) {
		switch i.(type) {
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Println("Wrong Type")
		}
	}
	mul(10)
	mul("100")
	mul(false)
}