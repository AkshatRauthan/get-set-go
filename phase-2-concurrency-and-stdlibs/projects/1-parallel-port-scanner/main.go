package main

import "fmt"

/*  REASON WHY LOCALHOST IS SUPER FAST

You send a "hello" to a closed port on your own machine (127.0.0.1).
Your computer immediately slams the door and yells, "Nobody is home, go away!"
This rejection message is the RST (Reset) packet.

Because your Go scanner heard this rejection instantly, it didn't wait around.
It immediately closed the connection and moved to the next port. That is how it scanned 65,535 ports in less than a second.
*/

/*  REASON WHY SERVERS ARE SUPER SLOW

You send a "hello" to a closed port on a remote server (like google.com). Instead of yelling "go away,"
the server's firewall just ignores you completely. Dead silence.

Because your Go scanner didn't get an RST packet, it just stood at the door waiting and waiting until your timeout finally expire.
*/

func main() {
	fmt.Print("\n|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|\n")
	fmt.Print("\nTesting Localhost first:\n")
	SequentialPortScanner("127.0.0.1", 1, 65535)

	//fmt.Print("\n\n|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|\n")
	//fmt.Print("\nTesting Nmap server now:\n")
	//SequentialPortScanner("scanme.nmap.org", 1, 200)

	fmt.Print("\n\n|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|\n")
	fmt.Print("\nTesting Localhost first:\n")
	ParallelPortScanner("127.0.0.1", 1, 65535, 5*1024)

	fmt.Print("\n\n|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|\n")
	fmt.Print("\nTesting Nmap server now:\n")
	ParallelPortScanner("scanme.nmap.org", 1, 65535, 5*1024)
}
