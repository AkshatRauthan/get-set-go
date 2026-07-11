package main

import (
	"fmt"
	"sync"
)

func CalculateFactorial2(n int, channel chan [2]int, wg *sync.WaitGroup) {
	defer wg.Done()

	ans := 1
	for i := 2; i <= n; i++ {
		ans *= i
	}

	// Because the channel is buffered, the worker drops the payload
	// into the buffer and immediately exits. It does not block.
	channel <- [2]int{n, ans}
}

func UsingWaitGroupsWithBufferedChannels() {
	// THE CHANGE: Add a buffer size (capacity) of 5.
	channel := make(chan [2]int, 5)

	var wg sync.WaitGroup

	wg.Add(5)

	go CalculateFactorial2(1, channel, &wg)
	go CalculateFactorial2(5, channel, &wg)
	go CalculateFactorial2(10, channel, &wg)
	go CalculateFactorial2(15, channel, &wg)
	go CalculateFactorial2(20, channel, &wg)

	// Since the buffer can hold all 5 results, the workers will not freeze.
	// Therefore, it is safe to block the main thread right here.
	wg.Wait()

	// Once wg.Wait() unblocks, we know the buffer is completely full.
	// We can safely close the channel now.
	close(channel)

	// The channel is closed, but the 5 results are safely sitting in the buffer!
	// The range loop will drain the buffer one by one and naturally exit.
	for result := range channel {
		fmt.Printf("Factorial of %d is %d\n", result[0], result[1])
	}
}
