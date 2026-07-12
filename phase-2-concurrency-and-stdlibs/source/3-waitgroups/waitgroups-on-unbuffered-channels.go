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

	/*
		Here, we are calling this part inside another goroutine because, we have used unbuffered channels.
		So every single goroutines last send statement is blocking in nature.
		[Cant get executed unless we are receiving in out for range loop]

		If we don't put the wait statement in the goroutine separately it will block our caller function, as we will
		not reach the for range loop as the wg.Wait() statement will block its execution.

		Now as we have place it inside a goroutine, it is safely detached from our caller function and our caller function
		will go on executing for range loop thus goroutines will complete their execution one by one.

		Here it may seem that the wg.Wait() statement is not needed as our for range statement wil itself program flow that
		is correct. But the main reason behind using it to trigger the close(channel) function at the right moment when all
		the goroutines are processed. So that our for range loop can terminate and we don't run into a deadlock.
	*/
	go func() {
		wg.Wait()      // Wait for all 5 workers to finish
		close(channel) // Safely close the channel ONLY when they are done
	}()

	for result := range channel {
		fmt.Printf("Factorial of %d is %d\n", result[0], result[1])
	}
}
