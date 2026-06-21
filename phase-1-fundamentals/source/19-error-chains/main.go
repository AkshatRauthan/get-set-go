package main

// RUN COMMAND:
// go run main.go error-chains.go

func main() {

	// 01. Sequential pipeline with explicit error propagation at every step
	// This directly replaces the TS pattern of chaining awaits inside try/catch
	// In Go: no implicit propagation — every call site handles or forwards the error manually
	ErrorChainsDemo()

	// 02. The broken %v version — shows how forgetting %w silently breaks errors.Is()
	ErrorChainBrokenDemo()
}
