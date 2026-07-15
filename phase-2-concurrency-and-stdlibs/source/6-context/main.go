package main

import "fmt"

/*
	Context: Immutable tree structures.
	Contexts in go are used primarily to bind all operations that are generated from a single parent function call
	together so that we can manage them properly.

	Context in Go is generally used to perform the following things:

	A. PROPER CANCELLATION AND GOROUTINE CLEANUP:
	   For stopping all the function calls in a call tree if the parent function returns and destroy all goroutines created inside them,
	   so that we can prevent memory leakage due to leaking goroutines.
	   Ex: If Func A returns, simply destroy all memory allocated to its child functions in call stack.

	B. TIMEOUTS AND DEADLINES:
	   For assigning a fixed central timeout/deadline for all operations/functions origination from the parent function of the call stack.
	   Ex:

	C. REQUEST-SCOPED VARIABLES:
	   For storing and passing request scoped variables like api-keys, auth-tokens, user-credentials etc. to all the underlying
	   functions of the call stack in backend systems.

	D. ABORT HANDLER:
	   For implementation of abort controller/handler pattern in backend systems.
	   [ Destroy all resources for an API Req. if requester disconnects ]
*/

func main() {
	fmt.Println("A. Context: Basics")
	ContextBasics()
}
