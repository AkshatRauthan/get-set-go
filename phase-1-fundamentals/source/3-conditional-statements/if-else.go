package main

import "fmt"

func IfElse() {
	var num int = -1

	fmt.Println("Enter input btw 1 and 10")
	fmt.Scanf("%d", &num)

	// 1. if else statements
	if num > 5 {
		fmt.Println("Number Greater than 5")
	} else if num >= 0 && num <= 5 {
		fmt.Println("Number btw 0 and 5")
	} else {
		fmt.Println("Invalid Number")
	}

	if age := 10; age > 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}
	// println(age) //Age cant be used out of scope from if else loop as its defined inside if else block
}
