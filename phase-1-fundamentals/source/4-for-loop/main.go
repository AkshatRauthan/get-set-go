package main

import "fmt"

func main() {
	var i, max int

	// Taking user input
	fmt.Printf("Enter starting point: ")
	fmt.Scanf("%d", &i) // Always use Scanf to take user input in Go
	fmt.Printf("Enter ending point: ")
	fmt.Scanf("%d", &max)

	// Go only has for loop but it can be used in diffrent ways....

	// 1. Usage As a while loop
	// for codition {
	//  }
	fmt.Println("\nUsing for loop as a While Loop")
	for i <= max {
		println(i)
		i++
	}

	fmt.Println("\nUsing for loop as a While Loop")
	// 2. Usage with range
	for x := range max {
		println(x)
	}

	// 3. fully fleged for loop like C++/Java
	// for ini; condi; incre/decre {
	// }
	for i := 0; i <= 2; i += 2 {
		println(i, ". Usage as a fully fleged for loop like in case of Java/C++")
	}

	// 4. Usage as infinite loop
	for {
		println("Terminate program with Ctrl + C")
	}

}
