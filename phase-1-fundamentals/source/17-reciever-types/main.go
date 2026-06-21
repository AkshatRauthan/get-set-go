package main

// RUN COMMAND:
// go run main.go receiver-types.go

func main() {

	// 01. Value receiver vs Pointer receiver — mutation difference
	ReceiverMutationDemo()

	// 02. Pointer receiver + interface satisfaction
	// The most important compile-time rule about receivers and interfaces
	ReceiverInterfaceDemo()
}
