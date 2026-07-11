package main

import "fmt"

// Buffered Channel:
// They are channels that can store a fixed number of instance of any variable/struct inside them as a buffer.
// They act as a fixed capacity [max size] queue where each send operation acts as a push and receive acts a remove.

// They are also blocking in nature, but they act differently then unbuffered channels:
// Here, the blocking behavior comes into play only when the requested operation is not possible on current buffer:
// - Func A performs read operation on empty channel, execution will block until a goroutine sends data onto channel
//   that func A can read.
// - Func B performs send operation on already filled channel, execution will block until someone reads data from channel
//	 and space is there for the new data. [Buffered channels are of fixed capacity]

// All other operations are non-blocking in nature:
// - Func A reading from non-empty channel.
// - Func B sending data to a non-full channel.

// Due to their above non-blocking behavior they are widely used in places like resource synchronization, producer-consumer
// patterns and as semaphores.

// When we have completed the use of a buffered channel we can turn it off using close() function.
// The main reason behind the usage of close is that if we are iterating over the buffer's data using
// for range statement then that loop will not get exited until we use the close function.

// IMPORTANT: There are some things are needed to be taken care of:
// A. NEVER CALL CLOSE ON A CHANNEL MORE THAN ONCE, CALLING CLOSE SECOND TIME WILL GIVE US AN INSTANT PANIC AND CRASH PROGRAM.
// B. To ensure above problem don't happen, call the CLOSE func on a channel only from sender's side so that it knows when to
//    stop sending the data.
// C. NEVER RECIEVE DATA FROM A CLOSED CHANNEL, AS IT WILL NOT THROW AN ERROR BUT WILL SILENTLY RETURN THE ZERO VALUE FOR ITS TYPE,
//	  THUS HIDING THE FACT THAT THE VALUE IS ACTUALLY STALE FOR US.
//    THEREFORE, USE THE `for range` STATEMENT WITH PROPER CLOSING OF CHANNEL, OR IF YOU ARE FETCHING SINGLE VALUE FROM BUFFER,
//    USE THE `val, ok : <- ch` CONSTRUCT SO THAT WE CAN CHECK IF CHANNEL IS OPEN OR NOT USING `ok` VALUE.

func fibonacci(n int, ch chan int64) {
	arr := [2]int64{0, 1}
	temp := int64(0)
	for i := 0; i <= n; i++ {
		if i < 2 {
			ch <- arr[i]
			continue
		}
		temp = arr[1] + arr[0]
		arr[0] = arr[1]
		arr[1] = temp
		ch <- temp
	}
	// Comment the below close statement to make the for range loop run infinitely and ran into a deadlock.
	close(ch)
}

func bufferedChannels() {
	n := 25
	ch := make(chan int64, 3)
	go fibonacci(n, ch)
	fmt.Print("The Fibonacci sequence up to F", n, " is : \n")
	for ele := range ch {
		fmt.Print(ele, " ")
	}

	// Here, channel is already in closed state [that's the reason above range func exited]
	fmt.Println("\nStale read from channel after closing:", <-ch)

	// Uncomment below line to get an instant PANIC as this is 2nd time close() called on channel ch
	//close(ch)
	fmt.Println()
}
