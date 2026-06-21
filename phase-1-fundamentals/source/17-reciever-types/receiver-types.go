package main

import "fmt"

/*
	Value Receiver vs Pointer Receiver:

	In Go, methods are defined with a receiver — the struct they belong to.
	The receiver can be a VALUE copy or a POINTER to the original.

	Value receiver  => func (o Order) MethodName()   => gets a COPY of the struct
	Pointer receiver => func (o *Order) MethodName() => gets the ACTUAL struct in memory

	This is the most important structural difference from TS classes, where methods
	always operate on the actual object (like pointer receivers).
	In Go, if you forget the *, your mutation silently disappears.

	RULE OF THUMB:
	- Method needs to mutate the struct     => pointer receiver (*Order)
	- Method only reads from the struct     => value receiver  (Order)
	- Struct is large                       => pointer receiver (avoids copying)
	- Implementing an interface             => be consistent across all methods
*/

type Order struct {
	id     int
	item   string
	amount float64
	status string
}

// Constructor pattern — same as your 11-structs example
func newOrder(id int, item string, amount float64) *Order {
	return &Order{
		id:     id,
		item:   item,
		amount: amount,
		status: "PENDING",
	}
}

// VALUE RECEIVER — receives a copy of Order, not the original
// Any changes made to `o` inside this method are discarded when it returns
// This is like receiving a TS object spread: { ...order } — a brand new copy
func (o Order) ApplyDiscountValue(pct float64) {
	// This mutates the LOCAL COPY only — the original Order is untouched
	o.amount = o.amount * (1 - pct/100)
	fmt.Printf("  Inside ApplyDiscountValue: amount = %.2f (local copy only)\n", o.amount)
}

// POINTER RECEIVER — receives a pointer to the actual Order in memory
// Changes made to `o` inside this method persist on the original struct
// This is equivalent to a TS class method mutating `this`
func (o *Order) ApplyDiscountPointer(pct float64) {
	// This mutates the ORIGINAL Order — the caller will see the change
	o.amount = o.amount * (1 - pct/100)
	fmt.Printf("  Inside ApplyDiscountPointer: amount = %.2f (original mutated)\n", o.amount)
}

// VALUE RECEIVER — safe for read-only operations, no mutation needed
func (o Order) Describe() string {
	return fmt.Sprintf("Order #%d | %s | $%.2f | %s", o.id, o.item, o.amount, o.status)
}

// POINTER RECEIVER — mutates status, must be pointer
func (o *Order) SetStatus(status string) {
	o.status = status
}

func ReceiverMutationDemo() {
	println("\n\n01. Value Receiver vs Pointer Receiver — Mutation:")

	order := newOrder(1, "Laptop", 1000.00)
	fmt.Println("Initial:", order.Describe())

	// Calling VALUE receiver method — original order.amount stays 1000
	fmt.Println("\nCalling ApplyDiscountValue(10):")
	order.ApplyDiscountValue(10)
	fmt.Println("After ApplyDiscountValue:", order.Describe()) // still $1000.00

	// Calling POINTER receiver method — original order.amount changes to 900
	fmt.Println("\nCalling ApplyDiscountPointer(10):")
	order.ApplyDiscountPointer(10)
	fmt.Println("After ApplyDiscountPointer:", order.Describe()) // now $900.00

	// SetStatus uses pointer receiver — mutation persists
	order.SetStatus("CONFIRMED")
	fmt.Println("\nAfter SetStatus('CONFIRMED'):", order.Describe())
}

/*
	Receiver Types and Interface Satisfaction:

	This is the most important compile-time rule to internalize:

	If a method is defined with a POINTER receiver (*Order),
	then ONLY *Order satisfies an interface requiring that method.
	A plain Order value does NOT satisfy it.

	If a method is defined with a VALUE receiver (Order),
	then BOTH Order and *Order satisfy the interface.

	Coming from TS: in TypeScript, a class instance always satisfies its interface
	regardless of how methods are defined. In Go, the receiver type is part of the
	method set — and it affects interface satisfaction at compile time.

	Your plan (page 4) calls this out as something to understand immediately.
*/

// Describer interface — any type with a Summary() method satisfies this
type Describer interface {
	Summary() string
}

type Product struct {
	name  string
	price float64
	stock int
}

// Summary defined with POINTER receiver — only *Product satisfies Describer
// Plain Product value does NOT have Summary() in its method set
func (p *Product) Summary() string {
	return fmt.Sprintf("Product: %s | Price: $%.2f | Stock: %d", p.name, p.price, p.stock)
}

// printDescription accepts any type satisfying Describer interface
func printDescription(d Describer) {
	fmt.Println(d.Summary())
}

func ReceiverInterfaceDemo() {
	println("\n\n02. Pointer Receiver + Interface Satisfaction:")

	// *Product satisfies Describer — pointer receiver, passing pointer => works
	p := &Product{name: "Mechanical Keyboard", price: 129.99, stock: 42}
	printDescription(p) // compiles fine

	// Plain Product value does NOT satisfy Describer because Summary() has pointer receiver
	// Uncommenting the lines below will cause a compile error:
	// "cannot use pVal (variable of type Product) as type Describer:
	//  Product does not implement Describer (Summary method has pointer receiver)"
	//
	// pVal := Product{name: "Mouse", price: 49.99, stock: 10}
	// printDescription(pVal) // COMPILE ERROR

	fmt.Println("\nOnly &Product (pointer) satisfies Describer — plain Product does not.")
	fmt.Println("This is because Summary() is defined with a pointer receiver (*Product).")
	fmt.Println("If Summary() had a value receiver (Product), both Product and *Product would satisfy Describer.")
}
