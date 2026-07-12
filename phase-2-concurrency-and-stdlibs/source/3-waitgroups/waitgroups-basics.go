package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

/*
	WaitGroups: WaitGroup is just a thread-safe counter (specifically, a type of semaphore).
	They are used as a counter that we use to ensure that all the goroutines spawned have completed
	their execution, and we don't terminate them prematurely.

	The value of a WaitGroup at any time tells us the number of uncompleted goroutines at that instance.

	They are implemented using hardware level atomic operations so they are guaranteed to be
	thread safe even at production level load.

	They have three methods:
	- A) wt.Add(n): Used to increment the value of waitgroup by n. n must always be +ive else it will panic.
	- B) wt.Done(): It is called when a goroutine completes its execution.
		  			This method decrements the waitgroup's value by 1.
	- C) wt.Wait(): This method blocks the execution of current thread/function until all the goroutines complete
					their work and the value of waitgroup becomes 0.

	Things to remember:
	- RULE OF ZERO: The counter starts at 0. Wait() only unblocks when the counter returns exactly to 0.
	- NEGATIVE PANIC: If the counter drops below zero, the WaitGroup immediately panics and crashes your program.
	- RACE CONDITION: You must always call Add() before you spawn the goroutine. If you spawn the goroutine first
	  and call Add() inside it, the main thread might reach Wait() before the goroutine has time to execute Add().
	  The WaitGroup will see the counter is 0 and let the main thread exit prematurely.

	Diffrence btw Waitgroups and Channels:
	- Channels -> Communication:
	  We use channels when goroutines need to talk to each other, pass data, or hand off state.
	  (Analogy: Passing a physical baton from one runner to another).
	- WaitGroups -> Coordination:
	  We use WaitGroups when we need to ensure that a batch of goroutines is finished.
	  (Analogy: A teacher counting heads on a bus to make sure all 30 students are back before telling the driver to leave).
*/

func NthVal(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Mimicking indefinite delay caused by processing.
	minWait, maxWait := 2, 10
	time.Sleep(time.Duration(rand.IntN(maxWait-minWait+1)+minWait) * time.Second)

	fmt.Println("Completed Process:", i)
}

// In this way we can make our function wait for the goroutines to execute fully before closing itself.
// [In all prior examples in ../2-channels and ../1-goroutines we ensured this behaviour using time.Sleep inside our caller function]

func WaitGroupsBasics() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go NthVal(1, &wg)
	go NthVal(2, &wg)

	wg.Wait()
}
