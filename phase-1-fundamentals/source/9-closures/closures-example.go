package main

import "fmt"

func createGreeting(greetingWord string) func(string) {
	greetPerson := func(personName string) {
		fmt.Println(greetingWord + "! " + personName)
	}
	return greetPerson
}

func greetPersons() {
	greetHello := createGreeting("Hello")
	greetHello("Akshat")
	greetHello("Lightning")

	greetGoodMorning := createGreeting("Good Morning")
	greetGoodMorning("Alan")
	greetGoodMorning("Aanand")

	createGreeting("Good Morning")("India")
	createGreeting("Hello")("Fuckerzz")
}
