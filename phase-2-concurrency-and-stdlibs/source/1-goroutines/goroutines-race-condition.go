package main

// Here in this file we will try to simulate a common problem while handling multiple goroutines.
// That is race condition during resource sharing amongst goroutines.

// In this function's output with -race flag we will get three kind of race conditions:
// a) Read vs Write — goroutine 21 reading while goroutine 28 is writing
// b) Write vs Write — goroutine 21 writing while goroutine 25 is writing simultaneously
// c) Write vs main goroutine read — goroutine 119 still writing while main already read the final value to print it

var count = 0

func incrementCounter() {
	count++
}

func goroutinesRaceCondition() {
	print("Initial count:", count)
	for i := 0; i < 100; i++ {
		go incrementCounter()
	}
	print("\nFinal Count after running 100 concurrent increment operations: ", count, "\n")
}
