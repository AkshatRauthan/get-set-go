package main

import "fmt"

/*
	04. Panic Basics:
	A panic is Go's mechanism for unrecoverable programmer errors — NOT for expected failures.

	Go itself panics on things like:
	- nil pointer dereference
	- index out of bounds
	- type assertion failure (without the ok idiom)

	When panic fires, the current function stops immediately, all its defers run,
	then the caller's defers run, all the way up the stack — until the program crashes
	with a stack trace, OR a recover() call intercepts it.

	CRITICAL RULE FROM YOUR PLAN (page 5):
	"Do not panic in library code."
	If a function is called by other code, return an error — never panic.
	Panic only for bugs that should never happen in correct usage
	(e.g. internal invariant violations in your own unexported logic).

	Coming from TS: panic is NOT throw. throw is for expected failures you handle.
	panic is for "this should be impossible and I want a loud crash with a stack trace."
	Expected failures in Go => return error. Impossible bugs => panic.
*/

func PanicBasics() {
	println("\n\n04. Panic Basics:")

	// Demonstrating a Go-triggered panic: nil map write
	// Uncommenting the lines below will crash the program with:
	// "assignment to entry in nil map"
	// var m map[string]int
	// m["key"] = 1 // Go panics here — you never called make()

	// Demonstrating a Go-triggered panic: index out of bounds
	// Uncommenting below crashes with: "index out of range [5] with length 3"
	// s := []int{1, 2, 3}
	// _ = s[5]

	// Manual panic: acceptable ONLY for truly impossible internal states
	// In a real trading system, this would be something like:
	// "order side is neither BUY nor SELL — this state should never reach here"
	fmt.Println("Panic left commented out intentionally — see comments above for examples.")
	fmt.Println("Run RecoverBasics() below to see panic interception via recover().")
}

/*
	05. Recover Basics:
	recover() stops a panicking goroutine from crashing the program.
	It ONLY works inside a deferred function — nowhere else.

	When to use recover():
	- At the TOP LEVEL of a server (e.g. per-request middleware) to prevent one bad request
	  from killing the whole server process.
	- In a goroutine wrapper that must not crash the parent process.

	When NOT to use recover():
	- As a replacement for returning errors from functions.
	- As a try/catch substitute — this is the TS habit to avoid.
	- In every function that might fail — only at process boundaries.

	The pattern is always:
	defer func() {
		if r := recover(); r != nil {
			// log it, return a safe value, continue
		}
	}()
*/

// riskyOperation simulates a function that panics on bad input
// In real code this might be a third-party library that panics unexpectedly
func riskyOperation(input int) {
	if input == 0 {
		panic("riskyOperation: received zero — internal invariant violated")
	}
	fmt.Println("riskyOperation completed successfully with input:", input)
}

// safeWrapper wraps riskyOperation and catches any panic via recover()
// This is the correct pattern: recover at the boundary, not inside every function
func safeWrapper(input int) (err string) {
	// defer with an anonymous function — recover() MUST be inside a deferred call
	defer func() {
		if r := recover(); r != nil {
			// r is the value passed to panic() — can be any type
			err = fmt.Sprintf("recovered from panic: %v", r)
		}
	}()

	riskyOperation(input)
	return "" // no error
}

func RecoverBasics() {
	println("\n\n05. Recover Basics:")

	// Case 1: normal input — no panic
	result := safeWrapper(5)
	if result != "" {
		fmt.Println("Error:", result)
	} else {
		fmt.Println("safeWrapper(5) completed without panic")
	}

	// Case 2: zero input — triggers panic inside riskyOperation, caught by recover in safeWrapper
	result = safeWrapper(0)
	if result != "" {
		fmt.Println("Caught:", result)
	} else {
		fmt.Println("safeWrapper(0) completed without panic")
	}

	fmt.Println("Program continues normally after recovered panic — this line always prints.")
}

/*
	06. Safe Divide Demo:
	This demonstrates recover() used to catch a divide-by-zero panic.
	BUT — read the comment inside carefully.
	This is shown so you recognize the pattern, NOT as a recommendation to use it.

	The correct Go way to handle division by zero is to return an error (like in 14-errors).
	recover() here is demonstrating the crash-containment use case ONLY.
*/

func divideInternally(a, b int) int {
	// integer division by zero is a Go runtime panic — not a returned error
	return a / b
}

func safeDivide(a, b int) (result int, err string) {
	defer func() {
		if r := recover(); r != nil {
			// intercepted the runtime panic — return a safe zero instead of crashing
			err = fmt.Sprintf("safeDivide recovered: %v", r)
			result = 0
		}
	}()

	result = divideInternally(a, b)
	return result, ""
}

func SafeDivideDemo() {
	println("\n\n06. Safe Divide Using Recover:")

	// NOTE: In production Go code, division by zero should be checked BEFORE dividing
	// and returned as an error — exactly like division() in 14-errors/basic-errors.go
	// recover() here is purely to demonstrate that Go runtime panics CAN be caught
	// You should NOT rely on this pattern as your primary error handling strategy

	res, err := safeDivide(10, 2)
	if err != "" {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("safeDivide(10, 2) =", res)
	}

	res, err = safeDivide(10, 0)
	if err != "" {
		fmt.Println("Caught runtime panic:", err)
	} else {
		fmt.Println("safeDivide(10, 0) =", res)
	}

	fmt.Println("Program still running after integer divide-by-zero panic was recovered.")
}
