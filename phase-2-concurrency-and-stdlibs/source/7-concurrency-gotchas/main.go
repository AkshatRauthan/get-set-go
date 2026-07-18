package main

import "fmt"

/*
	RACE CONDITIONS: where unprotected reads and writes overrides
	- Must besome data that is being written upon
	- could be a read-modify-write operation
	- and two or goroutines are working on it at same time
*/

/*
	DEADLOCKS: when no goroutine can make any progress
	- goroutines could all be blocked on empty channels
	- goroutines could all be blocked waiting on a mutex
	- GC could be prevented from running [busy loop]
	Go detects SOME deadlocks automatically; with -race it can find SOME data races too
*/

/*
	GOROUTINE LEAK: when parent process forgot to close goroutines and exits the scope
	- goroutines hangs up on a blocked or empty channel
	- not deadlock; other go routines progress
	- often found by looking at pprof output
	When you start a goroutine, always know how/when it will end
*/

/*
	CHANNEL ERRORS: unhandled channel errors or incorrect usage
	- trying to send on a closed channel
	- trying to send/recieve on a nil channel
	- closing a nil channel
	- closing on a channel twice
*/

/*
	OTHER CASES:
	- closure capture
	- misuse of Mutex
	- misuse of WaitGroup
	- misuse of select
*/

func main() {
	//fmt.Println("\nDeadlock Example 01:")
	//exampleDeadlock1()

	//fmt.Println("\nDeadlock Example 02:")
	//exampleDeadlock2()

	fmt.Println("\nDeadlock Due to Mutex example 01:")
	deadlockDueToMutex01()
}
