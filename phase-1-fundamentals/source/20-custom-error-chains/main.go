package main

// RUN COMMAND:
// go run main.go custom-error-chains.go

func main() {

	// 01. Custom error type propagated through a call chain, extracted with errors.As()
	// This is your TS `instanceof CustomError` pattern translated to Go
	CustomErrorChainDemo()

	// 02. The broken version — %v instead of %w silently kills errors.As()
	// The specific failure mode your TS propagation instinct won't warn you about
	CustomErrorChainBrokenDemo()
}
