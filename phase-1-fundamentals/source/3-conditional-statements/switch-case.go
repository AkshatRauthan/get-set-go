package main

import "fmt"

func switch_case() {
	i := 3
	// 2. switch case statements => Here no break needed in switch needed
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Default")
	}

	// Using Multiple cases in switch
	switch i {
	case 1, 3:
		fmt.Println("Odd")
	case 0, 2:
		fmt.Printf("Even")
	}
}
