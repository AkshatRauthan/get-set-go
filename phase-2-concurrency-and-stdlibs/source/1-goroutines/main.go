package main

// RUN COMMAND: go run .
// RUN COMMAND WITH RACE WARNINGS: go run -race .

func main() {
	println("\nA. Goroutines: Basics\n")
	goroutinesBasics()

	println("\n\nB. Goroutines: Race Contitions\n")
	goroutinesRaceCondition()
}
