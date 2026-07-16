package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func SequentialPortScanner(host string, start int, end int) {
	log.Printf("Starting sequential port scanning on host address %s from %d to %d\n\n", host, start, end)
	startTime := time.Now()

	for i := start; i <= end; i++ {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, i), 500*time.Millisecond)
		if err != nil {
			//fmt.Printf("Error connecting on %d\n", i)
			//fmt.Print(err.Error())
			continue
		}
		fmt.Printf("PORT %d OCCUPIED\n", i)
		_ = conn.Close()
	}

	timeTook := time.Since(startTime)
	seconds := int(timeTook.Seconds())
	milliseconds := int(timeTook.Milliseconds()) % 1000

	fmt.Print("\n")
	log.Print("Completed sequential port scanning\n")
	fmt.Printf("Time taken: %d seconds and %d milliseconds\n", seconds, milliseconds)
}
