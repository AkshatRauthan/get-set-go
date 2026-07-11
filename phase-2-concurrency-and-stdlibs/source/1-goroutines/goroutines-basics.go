package main

import "time"

// Goroutines are lightweight threads that are used to implement concurrency in Go.
// Instead of being actual OS threads they are built as a wrapper around them.
// Due to which they are very lightweight and fast to spawn.
// They are handled by the Go Scheduler instead of the OS.

// Concurrency: Working on several process efficiently at once by overlapping their execution period.
// Concurrency Do Not Mean Execution Of Multiple Tasks At Once. [That is parallelism of multiprocessing]

// Go routines are spawned by using the go keyword before the function call it needs to handle.

// Things to Remember:
// a) Irrespective of their spawning order, goroutines can be executed in any order....
// b) As soon as the main function returns, all goroutines will also be erased irrespective of their execution status....

// Output:
// Counter Finish:  7
// Counter Finish:  1
// Counter Finish:  4
// Counter Finish:  3
// Counter Finish:  8
// Counter Finish:  9
// Counter Finish:  5
// Counter Finish:  6
// Counter Finish:  2
// Counter Finish:  10

// In above output we can clearly see that their starting order as well as their completion order it totally random.
func spawnCounters(i int) {
	//println("Counter Init: ", i)
	time.Sleep(3 * time.Second)
	println("Counter Finish: ", i)
}

func goroutinesBasics() {
	for i := 1; i <= 10; i++ {
		go spawnCounters(i)
	}
	time.Sleep(5 * time.Second)
}
