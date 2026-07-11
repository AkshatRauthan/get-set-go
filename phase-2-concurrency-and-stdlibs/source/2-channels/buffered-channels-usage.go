package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

/*
	Usage Pattern 01: Worker pooling / Job queuing
	Distribute N jobs across M workers. Jobs sit in the buffer, workers pull from it as they finish.
*/

func Worker(jobPool chan int) {
	for jobId := range jobPool {
		minTimeout, maxTimeout := 1, 3
		currTimeout := rand.IntN(maxTimeout-minTimeout+1) + minTimeout

		// Simulating Processing Of Worker as Timeout
		time.Sleep(time.Duration(currTimeout) * time.Second)

		fmt.Println("Completed Job:", jobId)
	}
}

func WorkerPoolingPattern() {
	N, M := 15, 3

	// JobPool: We can process max 3 jobs at a time so we will load max 3 jobs.
	jobPool := make(chan int, M)

	// Spawning 3 Workers. [M = 3]
	for i := 1; i <= M; i++ {
		go Worker(jobPool)
	}

	// Simulated 15 Jobs. [N = 15]
	for i := 1; i <= N; i++ {
		jobPool <- i
	}
	fmt.Println("All jobs queued")

	time.Sleep(10 * time.Second)
	close(jobPool)
}

/*
	Usage Pattern 02: Usage As A Semaphore
	Limiting the maximum number of goroutines spawned by using buffered-channels as a semaphore.
*/

func TestProcess(i int, semaphore chan struct{}) {

	fmt.Println("Initiated Process:", i)
	time.Sleep(time.Second)
	//fmt.Println("Completed Process:", i)

	<-semaphore
}

func SemaphorePattern() {
	N, M := 15, 3 // N -> Total Processes, M -> Max allowed num of goroutines
	semaphore := make(chan struct{}, M)
	for i := 1; i <= N; i++ {
		semaphore <- struct{}{}
		go TestProcess(i, semaphore)
	}
	fmt.Println("All Processes Queued")
	time.Sleep(10 * time.Second)
}

/*
	Usage Pattern 03: Result Collection
	Fan-out to N goroutines, each sends result into a buffered channel sized N.
	Simply collection result of N goroutines into a buffered channel of size N. [Size N to ensure non-blocking nature]
*/

func ProcessJob(jobId int, results chan int) {
	// Simulating Timeout As Process Working => 1 or 2 sec sleep
	time.Sleep(time.Duration(rand.IntN(2)+1) * time.Second)
	results <- jobId
}

func ResultCollectionPattern() {
	N := 10
	results := make(chan int, N) // sized to N: So that all N goroutines send result without blocking

	for i := 1; i <= N; i++ {
		go ProcessJob(i, results)
	}

	for i := 1; i <= N; i++ {
		fmt.Println("Result Collected For JobId:", <-results)
	}

	fmt.Println("All Results Fetched")
}

/*
	Usage Pattern 04: Producer Consumer
	Producer generates faster than consumer processes. Buffered channel absorbs the burst — producer never blocks immediately.
	Consumer pulls at its own pace. Same idea as Kafka/RabbitMQ but in-process.
*/

func producer(ch chan int) {
	for i := 1; i <= 10; i++ {
		fmt.Println("Produced:", i)
		ch <- i
		// Producer is fast — no sleep
	}
	close(ch)
}

func consumer(ch chan int) {
	for val := range ch {
		// Consumer is slow — takes time to process each item
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Consumed:", val)
	}
}

func ProducerConsumerPattern() {
	// Buffer of 5 — producer can burst ahead by 5 items before blocking
	// Without buffer: producer blocks on every single item waiting for consumer
	ch := make(chan int, 5)

	go producer(ch)
	consumer(ch) // runs on main goroutine — blocks until channel closed and drained
}

// Main Entry Point
func bufferedChannelsUsagePatterns() {
	fmt.Println("Usage Pattern 01: Worker Pooling / Job Queuing")
	WorkerPoolingPattern()

	fmt.Println("\nUsage Pattern 02: Usage as a Semaphore")
	SemaphorePattern()

	fmt.Println("\nUsage Pattern 03: Result Collection Pattern")
	ResultCollectionPattern()

	fmt.Println("\nUsage Pattern 04: Result Collection Pattern")
	ProducerConsumerPattern()
}
