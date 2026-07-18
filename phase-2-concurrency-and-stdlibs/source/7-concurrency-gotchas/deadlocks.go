package main

import (
	"fmt"
	"sync"
	"time"
)

// Example will run find if flag = true.
// In case of flag = false, goroutine never sends onto bool channel, last read will cause deadlock
func exampleDeadlock1() {
	ch := make(chan bool)

	go func(flag bool) {
		fmt.Println("START")
		if flag {
			ch <- flag
		}
	}(false)
	<-ch
	fmt.Println("DONE")
}

// Again deadlock as we don't unlock mutex in 1st goroutine
func exampleDeadlock2() {
	mu := sync.Mutex{}
	done := make(chan bool)

	go func() {
		mu.Lock()
		//defer mu.Unlock()
	}()

	go func() {
		fmt.Println("START")
		time.Sleep(1 * time.Second)

		mu.Lock()
		defer mu.Unlock()

		fmt.Println("SIGNAl")
		done <- true
	}()

	<-done
	fmt.Println("DONE")
}
