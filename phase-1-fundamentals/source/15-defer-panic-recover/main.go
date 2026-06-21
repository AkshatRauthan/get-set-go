package main

// RUN COMMAND:
// go run main.go defer.go panic-recover.go

func main() {

	// 01. Defer Basics: what defer does, when it runs, and why it exists
	DeferBasics()

	// 02. Defer LIFO: multiple defers in one function run in reverse (Last In, First Out) order
	DeferLIFO()

	// 03. Defer With Early Return: defer fires even when a function exits early due to an error
	DeferWithEarlyReturn()

	// 04. Panic Basics: what panic is, when Go itself triggers it, and what NOT to use it for
	PanicBasics()

	// 05. Recover Basics: catching a panic before it crashes the whole program
	RecoverBasics()

	// 06. Safe Divide: using recover() as a crash-containment net — NOT as normal error handling
	SafeDivideDemo()
}
