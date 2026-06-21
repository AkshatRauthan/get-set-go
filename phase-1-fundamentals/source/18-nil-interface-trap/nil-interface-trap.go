package main

import "fmt"

/*
	The Nil Interface Trap:

	An interface value in Go is internally TWO things:
	  1. A TYPE descriptor  (what concrete type does this hold?)
	  2. A VALUE pointer    (where is the actual data?)

	An interface is nil ONLY when BOTH are nil.

	The trap: if you wrap a nil *ConcreteType inside an interface,
	the interface now has a non-nil TYPE descriptor (it knows it holds *ConcreteType)
	even though the VALUE pointer is nil.
	Result: the interface itself is NOT nil — even though the underlying pointer is.

	This is the #1 bug for developers coming from TS/JS, where null is just null.
	In Go, nil has TYPE information when it lives inside an interface.

	Coming from TS:
	In TypeScript, `null` and `undefined` carry no type information.
	if (error === null) always works predictably.
	In Go, a nil pointer wrapped in an interface is NOT nil — the type tag breaks equality.
*/

// AppError is a custom error type — same pattern as your 14-errors/20-custom-error-chains
type AppError struct {
	Code    int
	Message string
}

// Error() satisfies the built-in error interface
func (e *AppError) Error() string {
	return fmt.Sprintf("AppError [%d]: %s", e.Code, e.Message)
}

// THE BUGGY VERSION:
// This function intends to return "no error" by returning a nil *AppError.
// But it returns it as an `error` interface — and that wraps nil with type info.
// The caller checks `if err != nil` and gets TRUE — even though nothing went wrong.
func doWorkBuggy(succeed bool) error {
	var err *AppError = nil // nil pointer to AppError

	if !succeed {
		err = &AppError{Code: 500, Message: "something went wrong"}
	}

	// BUG: returning a typed nil pointer as an interface
	// Go wraps it in an error interface with type = *AppError, value = nil
	// The interface is NOT nil — it has a type descriptor
	return err
}

func NilInterfaceTrap() {
	println("\n\n01. The Nil Interface Trap:")

	// Case 1: succeed = true => doWorkBuggy returns a nil *AppError wrapped in error interface
	err := doWorkBuggy(true)

	// You'd expect this to print "No error" — but it prints "Got an error" instead
	// because the interface is non-nil (it carries the *AppError type tag)
	if err != nil {
		fmt.Println("Got an error:", err) // THIS FIRES even though we passed succeed=true
	} else {
		fmt.Println("No error — as expected")
	}

	// What the interface actually contains when doWorkBuggy(true) returns:
	// interface{ type: *AppError, value: nil }
	// The TYPE field is set => interface is NOT nil => err != nil is TRUE
	fmt.Println("\nWhy? Because the interface holds type=*AppError, value=nil")
	fmt.Println("An interface is nil ONLY when BOTH type AND value are nil.")
	fmt.Println("Wrapping a nil pointer inside an interface sets the type field — making it non-nil.")

	// Case 2: succeed = false => actual error, works as expected
	err = doWorkBuggy(false)
	if err != nil {
		fmt.Println("\nActual error case:", err.Error())
	}
}

/*
	The Fix:
	Never return a typed nil pointer as an interface.
	When your return type is `error` (an interface), return the untyped `nil` directly.
	Not `var e *AppError = nil; return e` — just `return nil`.

	`return nil` on an interface return type sets BOTH type and value to nil.
	That is a truly nil interface, and `err != nil` will correctly be false.
*/

// THE FIXED VERSION:
// Returns untyped nil directly when there is no error
// The caller's `err != nil` check now works correctly
func doWorkFixed(succeed bool) error {
	if !succeed {
		// Only construct the concrete type when there's an actual error
		return &AppError{Code: 500, Message: "something went wrong"}
	}

	// return nil directly — not a typed nil pointer — sets interface to truly nil
	return nil
}

func NilInterfaceFixed() {
	println("\n\n02. The Fix — Return Untyped nil Directly:")

	// Case 1: succeed = true => returns true nil interface => err != nil is false
	err := doWorkFixed(true)
	if err != nil {
		fmt.Println("Got an error:", err)
	} else {
		fmt.Println("No error — correctly nil this time") // THIS now fires correctly
	}

	// Case 2: succeed = false => actual error, works as expected
	err = doWorkFixed(false)
	if err != nil {
		fmt.Println("Actual error:", err.Error())
	}

	fmt.Println("\nKey rule: when returning an interface type (error), always return nil directly.")
	fmt.Println("Never return a typed nil pointer (var e *ConcreteType = nil; return e).")
}
