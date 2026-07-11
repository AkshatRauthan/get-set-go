package main

// CSP -> Communicating Sequential Processes =>
// "Do not communicate by sharing memory; instead, share memory by communicating."

// This is the philosophy on which the concurrency model of Go is based upon.

// Each single goroutine is an independently executing process handled by go scheduler instead of
// being directly bounded to an OS thread.

// Now these goroutines do not directly share memory between them. Instead, they use Channels.
// Channels are like connections that are utilized by goroutines to send/recieve data to/from other goroutines.

// By passing data back and forth, ownership of the data is cleanly handed off.
// Only one goroutine ever has access to a specific channel at one time,
// This is ensured by the blocking nature of channels.

// Channels also requires the goroutines to synchronize in order to use them.
// Ex. first A will send data into channel C only then B can fetch the data.
// If B already reached the fetching step but no data is in channel, it has to wait till A send the required data into channel.

// Now because no resources are directly being shared amongst goroutines traditional problems like race conditions, deadlocks
// Don't occur in go until we are using Channels solely amongst goroutines.

// If No. of read operations on a channel > No. of write operations on a channel:
// - Flow Stops completely as no data to be received
// - So program run into deadlock.

// If No. of read operations on a channel < No. of write operation on a channel:
// - The program will exit before all the goroutines complete their execution.

func sum(arr []int, c chan int) {
	sum := 0
	for _, val := range arr {
		sum += val
	}

	// Sending data to channel from a goroutine
	c <- sum
}

func channelsBasics() {
	// Defining a new channel
	c := make(chan int)
	s := []int{1, 4, -10, 11, 5, 6, 3, 2, -2, -5}

	go sum(s[:len(s)/2], c) // 11
	go sum(s[len(s)/2:], c) // 4

	// Receiving data from a channel via a goroutine
	x, y := <-c, <-c
	// Either 11, 4 OR 4, 11

	print(x, " ", y, " ", x+y, "\n\n")
}
