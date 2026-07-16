package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func ParallelPortScanner(host string, start int, end int, limit int) {
	log.Printf("Starting parallel port scanning on host address %s from %d to %d\n", host, start, end)
	startTime := time.Now()

	limiter := make(chan struct{}, limit)
	defer close(limiter)

	wg := &sync.WaitGroup{}
	wg.Add(end - start + 1)

	for i := start; i <= end; i++ {

		go func(i int) {
			defer wg.Done()
			limiter <- struct{}{}
			defer func() { <-limiter }()
			conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, i), 1000*time.Millisecond)
			if err != nil {
				//fmt.Printf("Error connecting on %d\n", i)
				//fmt.Print(err.Error())
				return
			}
			fmt.Printf("PORT %d OCCUPIED\n", i)
			_ = conn.Close()
		}(i)
	}

	wg.Wait()
	timeTook := time.Since(startTime)
	seconds := int(timeTook.Seconds())
	milliseconds := int(timeTook.Milliseconds()) % 1000

	fmt.Print("\n")
	log.Print("Completed parallel port scanning\n")
	fmt.Printf("Time taken: %d seconds and %d milliseconds\n", seconds, milliseconds)
}
