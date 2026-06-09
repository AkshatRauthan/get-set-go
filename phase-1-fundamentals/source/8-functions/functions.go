package main

import "fmt"

// defining a normal function
func add(a int, b int) int {
	return a+b
}

// function returning multiple values
func calc(a, b int, c,d string) (int, int, int, float32, string) {
	return a+b, a-b, a*b, float32(a/b), c+d
}

func functions() {
	fmt.Println("\nUsing Functions In Go:")

	// calling function with mul return types....
	r1, r2, r3, r4, r5 := calc(1,1,"hello", "world")
	fmt.Println("r1: ",r1) // 2
	fmt.Println("r2: ",r2) // 0
	fmt.Println("r3: ",r3) // 1
	fmt.Println("r4: ",r4) // 1
	fmt.Println("r5: ",r5) // helloworld

	// calling functions with mul. return types ignoring some values....

	r1, r2, r3, r4, _ = calc(2, 2, "alan", "walker")
}
