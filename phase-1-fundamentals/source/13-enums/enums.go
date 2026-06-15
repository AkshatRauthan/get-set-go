package main

type OrderStatus int

const (
	RECIEVED OrderStatus = iota // auto numbers from 0 onwards...
	CONFIRMED
	DELAYED
	CANCELLED
)

type MessageString string

const (
	SIE  = "Congratulations! Your request is accepted."
	NIEN = "Request Processing! Please be patient while we review your request."
)

func enums() {
	println("\nGo doesnt have built-in support for enums....")
	println("We define them explicitly by following the above workaround....")

	println("\nUsing Enums:\nOrder Status:", RECIEVED)
	println("Order String: ", SIE)
}
