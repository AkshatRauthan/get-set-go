package main

import "fmt"

func main() {
	fmt.Println("\nA. Channels: Basic\n")
	channelsBasics()

	fmt.Println("\nB. Unbuffered Channels\n")
	unbufferedChannels()

	fmt.Println("\nC. Buffered Channels\n")
	bufferedChannels()

	fmt.Println("\nD. Unbuffered Channels: Usage Patterns\n")
	unbufferedChannelsUsagePatters()

	fmt.Println("\n\nE. Buffered Channels: Usage Patterns\n")
	bufferedChannelsUsagePatterns()
}
