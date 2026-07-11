package main

import "fmt"

func main() {
	// variables
	var name string = "developer-akshat"
	var age uint = 21 // unsigned integers
	var isDeveloper bool = true

	//without type declaration
	var c64 = 5 + 7i // complex numbers with 32-bit real and imaginary parts => (5+7i)

	c64 = complex(10, imag(c64)) // updating real part
	c64 = complex(real(c64), 15) // updating imaginary part

	// constants
	const pie float32 = 3.14

	// multiple variable declaration
	const (
		a string = "a"
		b string = "b"
	)
	const c, d int = 1, 2

	// shorthand declaration
	x, y := 10, 20
	shortVar := "I am a short declaration string"

	x += y
	shortVar += " and I can be modified"
	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
	fmt.Println("isDeveloper: ", isDeveloper)
	fmt.Println("pie: ", pie)
	fmt.Println("c64: ", c64)
	fmt.Println("  ")

	// In shorthand declaration if we are declaring/initializing multiple variables....
	// Then using := will work only if there is at least one new variable in left side....
	// x, y := 20, 10     ✖ [both x and y are already initialized]
	// x,y,z := 10,20,30  ✔ [z is newly initialized]

	// We cannot declare global (package-level) variables in Go using the shorthand declaration operator (:=).
	// The shorthand syntax is strictly limited to function bodies and will cause a compilation error if used outside a function.
	// According to the Go language specification, every construct outside a function must begin with a distinct keyword
	// (such as package, import, var, const, type, or func). As := is basically an assignment operator it will threw compiler error.
}
