package main

import "fmt"

func main() {
	fmt.Println("A. Using WaitGroups With Unbuffered Channels:")
	UsingWaitGroupsWithUnbufferedChannels()

	fmt.Println("\nB. Using WaitGroups With Buffered Channels:")
	UsingWaitGroupsWithBufferedChannels()
}
