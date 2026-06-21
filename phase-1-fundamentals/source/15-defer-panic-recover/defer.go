package main

import (
	"errors"
	"fmt"
)

/*
	01. Defer Basics:
	defer schedules a function call to run AFTER the surrounding function returns — no matter how it returns.
	This is Go's primary cleanup mechanism, replacing C++'s RAII destructors and TS's try/finally blocks.

	The most important rule: defer the cleanup IMMEDIATELY after acquiring the resource.
	Not at the bottom of the function — right after the open/acquire call.
	This way, even if you add an early return later, cleanup is guaranteed.

	Think of it as: "I just opened this — I promise to close it when I'm done."
*/

// openDBConnection simulates acquiring a resource (like a DB connection or file handle)
func openDBConnection(name string) string {
	fmt.Println("Opened connection to:", name)
	return name
}

func closeDBConnection(name string) {
	fmt.Println("Closed connection to:", name)
}

func DeferBasics() {
	println("\n\n01. Defer Basics:")

	conn := openDBConnection("orders-db")

	// defer closeDBConnection immediately after opening — not at the bottom
	// This guarantees cleanup even if the function returns early due to error below
	defer closeDBConnection(conn)

	// simulating work that could fail
	fmt.Println("Running query on:", conn)
	fmt.Println("Query done.")

	// closeDBConnection will run here automatically — after this function returns
	// You will see "Closed connection to: orders-db" printed AFTER "Query done."
}

/*
	02. Defer LIFO (Last In, First Out):
	When multiple defers exist in one function, they execute in reverse order of how they were registered.
	This mirrors how you'd manually unwind a stack of acquired resources:
	last acquired => first released.

	Example: Open DB => Open Cache => Open Logger
	Defer order: Logger closes first, then Cache, then DB
*/

func DeferLIFO() {
	println("\n\n02. Defer LIFO Ordering:")

	// Each defer is registered in order — but they fire in reverse
	defer fmt.Println("Step 3 deferred — fires FIRST (last registered)")
	defer fmt.Println("Step 2 deferred — fires SECOND")
	defer fmt.Println("Step 1 deferred — fires LAST (first registered)")

	fmt.Println("Function body executing...")
	fmt.Println("All defers registered. Function about to return.")

	// Output order when this returns:
	// "All defers registered. Function about to return."
	// "Step 3 deferred — fires FIRST"
	// "Step 2 deferred — fires SECOND"
	// "Step 1 deferred — fires LAST"
}

/*
	03. Defer With Early Return:
	This is the most practically important defer behaviour.
	In TS you'd use try/finally to guarantee cleanup on early exits.
	In Go, defer replaces finally entirely — it fires on ANY return path:
	normal return, early return, or even a panic.

	Key insight: if you're coming from TS's try/finally pattern,
	defer is your direct replacement — cleaner, and always correct if placed right after resource acquisition.
*/

var ErrInvalidOrderID = errors.New("Error: Order ID must be positive")
var ErrOrderNotFound = errors.New("Error: Order not found in system")

func processOrder(orderID int) error {
	fmt.Println("\nOpening order processor for orderID:", orderID)

	// defer cleanup immediately — before any logic that could return early
	defer fmt.Println("Order processor closed for orderID:", orderID)

	// early return path 1: invalid input
	if orderID <= 0 {
		return ErrInvalidOrderID // defer still fires here
	}

	// simulated order DB
	orders := map[int]string{1: "Laptop", 2: "Monitor"}
	order, ok := orders[orderID]

	// early return path 2: not found
	if !ok {
		return ErrOrderNotFound // defer still fires here too
	}

	// happy path
	fmt.Println("Processing order:", order)
	return nil // defer fires here as well
}

func DeferWithEarlyReturn() {
	println("\n\n03. Defer With Early Return:")

	// Case 1: invalid ID — hits early return path 1
	err := processOrder(-1)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	// Case 2: valid ID, order not found — hits early return path 2
	err = processOrder(99)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	// Case 3: valid ID, order exists — hits happy path
	err = processOrder(1)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	// In every case above, "Order processor closed for orderID: X" prints
	// regardless of which return path was taken — that is the guarantee defer gives you
}
