package main

import (
	"fmt"
	"sync"
)

func CalculateFactorial(n int, channel chan [2]int, wg *sync.WaitGroup) {
	defer wg.Done()

	ans := 1
	for i := 2; i <= n; i++ {
		ans *= i
	}
	channel <- [2]int{n, ans}
}

func UsingWaitGroupsWithUnbufferedChannels() {
	channel := make(chan [2]int)
	wg := sync.WaitGroup{}

	wg.Add(5) // We will spawn 5 goroutines so we will initialise the WaitGroup with value 5

	go CalculateFactorial(1, channel, &wg)
	go CalculateFactorial(5, channel, &wg)
	go CalculateFactorial(10, channel, &wg)
	go CalculateFactorial(15, channel, &wg)
	go CalculateFactorial(20, channel, &wg)

	go func() {
		wg.Wait()      // Wait for all 5 workers to finish
		close(channel) // Safely close the channel ONLY when they are done
	}()

	for result := range channel {
		fmt.Printf("Factorial of %d is %d\n", result[0], result[1])
	}
}
