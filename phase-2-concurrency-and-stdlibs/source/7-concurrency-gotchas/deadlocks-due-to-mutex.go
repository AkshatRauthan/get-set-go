package main

import (
	"fmt"
	"sync"
	"time"
)

// Deadlock dut to multiple mutex usage:
// Deadlock: both goroutines want lock on both 1 and 2 mutex.
// 1st one locks 1, 2nd one locks 2nd  one, no one bot both deadlock.

// In case of more than one mutex, always keep locking order of mutex same across all goroutines.
func deadlockDueToMutex01() {
	mutex1 := sync.Mutex{}
	mutex2 := sync.Mutex{}

	done := make(chan struct{})
	fmt.Println("START")

	go func() {
		mutex1.Lock()
		defer mutex1.Unlock()

		time.Sleep(1 * time.Second)

		mutex2.Lock()
		defer mutex2.Unlock()

		fmt.Println("SIGNAL")
		done <- struct{}{}
	}()

	go func() {
		// Wrong snippet

		//mutex2.Lock()
		//defer mutex2.Unlock()
		//
		//time.Sleep(1 * time.Second)
		//
		//mutex1.Lock()
		//defer mutex1.Unlock()

		// Correct Snippet

		mutex1.Lock()
		defer mutex1.Unlock()

		time.Sleep(1 * time.Second)

		mutex2.Lock()
		defer mutex2.Unlock()

		fmt.Println("SIGNAL")
		done <- struct{}{}
	}()

	<-done
	fmt.Println("DONE")
	<-done
	fmt.Println("DONE")
}
