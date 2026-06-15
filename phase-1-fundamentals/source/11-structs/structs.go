package main

import (
	"fmt"
	"time"
)

type Order struct { // below all are optional will append zero values of fields by default
	id        string
	amount    int
	status    string
	createdAt time.Time
}

// GO structs dont have contructors, we follow this convention to define a constructor using a function.....
func newOrder(id string, amount int, status string) *Order {
	myOrder := Order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

// Defining struct methods:
// Structs in GO are pass by value, therefore we need to take a Pointer to the struct object.
// We don't need to use dereference operators with structs thet dereference automatically in GO....
func (o *Order) setStatus(status string) {
	println("Setting status of struct using struct method....")
	o.status = status
}
func (o Order) getStatus() string {
	println("Getting status of struct using struct method....")
	return o.status
}

func structs() {
	var order01 = Order{
		id:     "1",
		amount: 10,
	}
	order01.status = "Done"
	fmt.Println("Struct: ", order01)
}
