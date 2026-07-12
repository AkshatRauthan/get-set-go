package main

import "fmt"

func main() {
	fmt.Println("A. WaitGroups: Basics")
	WaitGroupsBasics()

	fmt.Println("\nB. Using WaitGroups With Unbuffered Channels:")
	UsingWaitGroupsWithUnbufferedChannels()

	fmt.Println("\nC. Using WaitGroups With Buffered Channels:")
	UsingWaitGroupsWithBufferedChannels()
}
