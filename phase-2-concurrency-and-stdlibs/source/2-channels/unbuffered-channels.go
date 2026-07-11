package main

import "time"

// Unbuffered Channels:
// They are basic channels that can share a single variable/struct instance at a time.

// They are blocking in nature:
// - If func A sends data to an unbuffered channel, the flow of A will stop at that statement until another goroutine reads that data.
// - If func B reads data from an unbuffered channel, the flow of func B will stop until it receives that data from channel. And in that case
//	 if the channel currently is empty then B's flow will stop until another goroutine send data to the channel.

// Due to the above property of unbuffered channel to hinder func execution it is widely used as a mutex or for synchronization.

type s = struct{}

// Always use this type for signals as struct{} objects have 0 bit of allocation

func testUnbufferedChannel(i int, ch chan s) {
	//print("Starting Execution Of Test Function Number ", i, "\n")
	time.Sleep(2 * time.Second)
	var signal s
	ch <- signal
	print("Completed Execution Of Test Function Number ", i, "\n")
}

func unbufferedChannels() {
	ch := make(chan s)

	for i := 0; i <= 3; i++ {
		//print("Called Test Function Number ", i, "\n")
		go testUnbufferedChannel(i, ch)
		//<-c // Uncomment this statement for sequential execution of goroutines
	}
	// Uncomment below for loop for parallel execution of goroutines
	for i := 0; i <= 3; i++ {
		<-ch
	}
	time.Sleep(10 * time.Millisecond)
	print("Completed The Execution Of Test Functions", "\n")
}
