package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	PROPAGATION OF CONTEXT:
	- The parent functions generates the root context also known as background context.
	- Then it generates its own context using based upon the background context, and passes this new context to its children.
	- Now context is immutable in nature, so the child function creates its own context based upon the context of
	  it received. So that cancelling the parent can also lead to the cancellation of the child.
	- The child then passes this context to its own children and this keeps going on.

	[We can add indivisual timeout/deadline to the child context's but its value must be smaller than parent context's value]

	Context provides us two things:
	- A Done channel that closes when the cancellation occurs. [Implements an internal done channel so that we can close goroutines]
	- An error value that's readable once the channel closes. [Tells weather the request was cacelled or timed out]

	Using a Channel:
	ctx, cancel := contex.func()
	Here, ctx is the context object to be passed on to the childrens
		  cancel is the function we will use to trigger cancellation of our context [meant to be used only inside function scope where defined]
		  func() is the respective function we are using to create a context based on our needs.

	Diffrent types of contexts constructors....
	- context.Background(): 					Root/Background Context  		[Don't return a cancel function]

	- context.WithCancel(parent): 				Cancellation by Cancel only
	- context.WithDeadline(parent, time): 		Cancellation by Cabcel + Auto-cancel at given time
	- context.WithTimeout(parent, timeout):		Cancellation by Cancel + Auto-cancel after timeout
	- context.WithValue(parent, key, value):   	Cancellation by Cancel only + Req-Scoped Variables 	[Don't return a cancel function]

	THINGS TO REMEMBER:
	- The root function of the call stack don't need to create a new context. They can simply use ctx.Done() channel.
	- Always pass context as the FIRST argument to any function that do I/O.
	- Always defer cancel() immediately after creating a cancellable context.
	  Forgetting cancel() = goroutine leak — the goroutine lives forever waiting on ctx.Done()

	We can also chain contexts inside a single function to use a combination of above things like in the snippet below...

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ctx = context.WithValue(ctx, "requestID", "abc-123")
	ctx = context.WithValue(ctx, "userID", 42)

	- Backgorund() and WithValue() dont return cancel values.
	  As background channel will be overwritten by another channel and WithValue() is always chained onto another constructor.
*/

func Ticker(min, max int) (<-chan struct{}, func()) {
	tickerChan := make(chan struct{})
	stopChan := make(chan struct{})

	go func() {
		// Ensure the channel is closed when the goroutine exits
		defer close(tickerChan)
		defer close(stopChan)

		delta := max - min
		initialWait := min + delta
		timer := time.NewTimer(time.Duration(initialWait) * time.Second)
		defer timer.Stop()

		for {
			select {
			case <-timer.C:
				// Non-blocking send due to default case of internal select statement
				select {
				case tickerChan <- struct{}{}:
				default:
				}

				randomWait := min + rand.Intn(delta)
				timer.Reset(time.Duration(randomWait) * time.Second)

			case <-stopChan:
				return
			}
		}
	}()

	s := sync.Once{}
	closeTicker := func() {
		s.Do(func() {
			stopChan <- struct{}{}
		})
	}

	return tickerChan, closeTicker
}

func GenericStreamSender(ctx context.Context, wg *sync.WaitGroup, dataItem string, stream chan<- string, N int) {
	defer wg.Done()

	// Attaching a ticker that sends ticks at random interval
	ticker, closeTicker := Ticker(0, 3)
	defer closeTicker()

	// Attaching a timeout for automated closure of the goroutine: btw 5 to 25 sec
	rootTimeout := rand.Intn(21) + 5
	rootTimer := time.After(time.Duration(rootTimeout) * time.Second)

	// If tick received push to stream
	for {
		select {
		case <-ticker:
			stream <- fmt.Sprintf("%s: %d", dataItem, N)
		case <-ctx.Done():
			fmt.Printf("Done Signal On Sender %d For %s\n", N, dataItem)
			return
		case <-rootTimer:
			fmt.Printf("Auto Timeout On Sender %d For %s\n", N, dataItem)
			return
		}
	}
}

func GenericStreamReader(ctx context.Context, dataItem string, stream <-chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Done Signal On Reader For %s\n", dataItem)
			return
		case data, ok := <-stream:
			if !ok {
				return
			}
			fmt.Println(data)
		}
	}
}

func InfiniteStreamGenerator(ctx context.Context, mainWg *sync.WaitGroup, dataItem string, N int) {
	defer fmt.Printf("Closed Infinite Stream Generator For %s\n", dataItem)
	defer mainWg.Done()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// WaitGroups for internal generators. [Exit function only after all generators stop]
	wg := &sync.WaitGroup{}
	wg.Add(N)

	// Unbuffered channel for Non blocking behaviour
	dataStream := make(chan string, N)

	// Spawning a single reader for a single data item.
	go GenericStreamReader(ctx, dataItem, dataStream)

	// Spawning N generators per data item.
	for i := 1; i <= N; i++ {
		go GenericStreamSender(ctx, wg, dataItem, dataStream, i)
	}

	// Waiting for all generators to stop before exiting
	wg.Wait()
	close(dataStream)
}

func ContextBasics() {
	wg := &sync.WaitGroup{}
	wg.Add(3)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go InfiniteStreamGenerator(ctx, wg, "Respect", 3)
	go InfiniteStreamGenerator(ctx, wg, "Power", 3)
	go InfiniteStreamGenerator(ctx, wg, "Banana", 3)

	wg.Wait()
}
