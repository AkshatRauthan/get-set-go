package main

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

/*
	SELECT STATEMENT: The select statement allows us to either send data to a channel or receive data from a channel
	as soon as they come in ready state. We can pass multiple channels on to the select statement just like we do in switch
	statement.

	The SELECT statement is generally used inside a LOOP, that loop allows us to process channels as soon as they come in
	READY state and then again start listening to all the channels.
*/

/*
	Why we need select statement?
	Because when we try to read or send data from multiple channels from a same part we are bounded by their wait times until they come in Ready state.
	Assume that if the first channel that we are reading is slow and until it comes into ready state, there are some other channels that reached
	ready state, nut they are placed below, then they have to wait until all channels placed above them in loop finish their operations.

	To remove this behaviour, we use select statement. It directly executes the channel that comes in ready state first.
*/

/*
	If at any time no channel is in ready state it executes the default case if there is any.
	However, we must beware of using default statement especially when we are using SELECT WITH LOOPS.
	This is because if we use select statement inside a loop and the default case is there, it may run huge no of times as if channel
	operations are not that frequent that inturn can block our CPU by consuming a large part of its processing.

	SO NEVER USE Print or Log statements inside DEFAULT CASE of a SELECT statement inside a LOOP.
*/

/*
	IN MY SYSTEM IF I RUN THE BELOW PROGRAM FOR 10 SECONDS GUESS HOW MANY TIMES THE DEFAULT CASE GETS EXECUTED:
	Default Case Executed 58045405 times in 10 seconds

	AND IT LOCKED THE CPU THREAD ON WHICH IT IS RUNNING AT 100% USAGE.
	SO ALWAYS TRY TO AVOID DEFAULT CASE INSIDE SELECT STATEMENTS THAT ARE RUNNING INSIDE A LOOP. [ANY LOOP BE IT DEFAULT OR WITH LIMIT]
*/

func SelectStatementBasics() {
	N := 10 // Program run duration in Sec.

	counter := atomic.Uint64{}

	// Initialising multiple channels
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	// Pushing data periodically inside these channels
	for i := range chans {
		go func(i int, ch chan<- int) { // Using a write-only channel
			for { // Infinite loop for pushing data into channel
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}
		}(i+1, chans[i])
	}
	// Looping over those channels using select statement
	fmt.Println("Stating Select Statement")

	// Stopper: It is a channel that returns a result after N seconds
	// So we can listen up to 30 seconds.
	stopper := time.After(time.Duration(N) * time.Second)

	// Using a named infinite loop so that we can directly break if inside select statement.
	// If we directly use break without loop name it break will just leave the select statement and continue the for loop.
loop:
	for {
		select {
		case c1 := <-chans[0]:
			log.Println("Pinged:", c1)
		case c2 := <-chans[1]:
			log.Println("Pinged:", c2)
		case <-stopper:
			log.Println("Program Stopper Encountered")
			break loop
			//default: // Uncomment this part to see why it is fatal to run DEFAULT IN LOOP
			//	counter.Add(1)
		}
	}
	fmt.Println("Closed Listeners")
	fmt.Printf("Default Case Executed %d times in %d seconds\n", counter.Load(), N)
}
