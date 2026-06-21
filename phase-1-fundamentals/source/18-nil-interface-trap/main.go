package main

// RUN COMMAND:
// go run main.go nil-interface-trap.go

func main() {

	// 01. The nil interface trap — the single most common bug for new Go developers
	// Your plan (page 4) flags this explicitly: a nil pointer wrapped in an interface is NOT nil
	NilInterfaceTrap()

	// 02. The correct fix — how to actually return a nil error safely
	NilInterfaceFixed()
}
