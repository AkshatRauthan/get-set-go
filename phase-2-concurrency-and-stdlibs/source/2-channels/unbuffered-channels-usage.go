package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// Usage Pattern 01: Done signals / goroutine synchronisation
// "Tell me when you're finished" — Already implemented in ./unbuffered-channels.go as testUnbufferedChannels()

// Usage Pattern 02: Request-response between two goroutines
// One goroutine sends a request, waits. Other goroutine processes, sends response back.

type Request struct {
	price   int
	channel chan Response
}
type Response struct {
	price         int
	success       bool
	transactionId int
	error         string
}

func processPayment(req Request) {
	fmt.Println("Processing Payment request...")

	// Simulating real life transaction processing....
	time.Sleep(5 * time.Second)

	// Sending Response Back For The Request
	minId, maxId := 10000000, 99999999
	transactionId := rand.IntN(maxId-minId+1) + minId
	success := rand.IntN(2) == 1
	err := ""
	res := Response{
		price:         req.price,
		success:       success,
		transactionId: transactionId,
		error:         err,
	}
	req.channel <- res

	// Simulating some other tasks that we don't want our below func to wait for
	time.Sleep(time.Second)
	fmt.Println("Completed Full Payment Processing in helper function")
}

func requestResponsePattern() {
	ch := make(chan Response)
	req := Request{
		price:   rand.IntN(2),
		channel: ch,
	}
	go processPayment(req)

	res := <-ch
	if res.success {
		fmt.Println("Payment processed successfully")
	} else {
		res.error = "REQUEST_TIMED_OUT"
		fmt.Println("Payment failed with error: ", res.error)
	}
	fmt.Println("Transaction ID: ", res.transactionId)
	time.Sleep(time.Second + time.Millisecond)
}

// Usage Pattern 03: Guaranteed handoff
// When you need certainty that the receiver actually got the value before the sender continues.

type Order struct {
	orderId int
	symbol  string
	metricA int
	metricB int
	metricC int
	channel chan OrderAnalysis
}

type OrderAnalysis struct {
	orderId    int
	symbol     string
	isSafe     bool
	riskStatus string
}

func riskAnalyser(order Order) {
	a, b, c := order.metricA, order.metricB, order.metricC
	normalisedScore := (8*a + 6*b + 6*c) / (20)

	isSafe := false
	riskStatus := ""

	if (normalisedScore >= 80) || (normalisedScore >= 70 && a >= 90) {
		riskStatus = "SAFE"
		isSafe = true
	} else if normalisedScore >= 70 {
		riskStatus = "NORMAL"
		isSafe = true
	} else if normalisedScore >= 60 && a >= 80 {
		riskStatus = "MODERATE"
		isSafe = true
	} else if normalisedScore >= 50 {
		riskStatus = "HIGH"
	} else {
		riskStatus = "VERY_HIGH"
	}

	analysis := OrderAnalysis{
		isSafe:     isSafe,
		orderId:    order.orderId,
		symbol:     order.symbol,
		riskStatus: riskStatus,
	}

	order.channel <- analysis
	fmt.Println("Risk Assessment Completed For Symbol: ", analysis.symbol)
}
func getNewOrder(symbol string, channel chan OrderAnalysis) Order {
	minId, maxId := 10000000, 99999999
	order := Order{
		orderId: rand.IntN(maxId-minId+1) + minId, symbol: symbol, metricA: rand.IntN(101), metricB: rand.IntN(101), metricC: rand.IntN(101), channel: channel,
	}
	return order
}

func guaranteedHandoffPattern() {
	channel := make(chan OrderAnalysis)

	order1, order2, order3 := getNewOrder("BTC", channel), getNewOrder("ETH", channel), getNewOrder("IDEA", channel)
	order4, order5 := getNewOrder("BANK_NIFTY", channel), getNewOrder("MRF", channel)
	go riskAnalyser(order1)
	go riskAnalyser(order2)
	go riskAnalyser(order3)
	go riskAnalyser(order4)
	go riskAnalyser(order5)

	for i := 0; i < 5; i++ {
		analysis := <-channel
		if analysis.isSafe {
			fmt.Println("Symbol", analysis.symbol, "Classified As Safe To Trade")
		} else {
			fmt.Println("Symbol", analysis.symbol, "Classified As Unsafe To Trade")
		}
		fmt.Println("Risk Status: ", analysis.riskStatus)
	}
	time.Sleep(100 * time.Millisecond)
}

// Entry point
func unbufferedChannelsUsagePatters() {
	fmt.Println("Usage Pattern 01: Done signals / goroutine synchronisation")
	fmt.Println("Already implemented as testUnbufferedChannels()")

	fmt.Println("\nUsage Pattern 02: Request-response between two goroutines")
	requestResponsePattern()

	fmt.Println("\nUsage Pattern 03: Guaranteed handoff")
	guaranteedHandoffPattern()
}
