package main

// Interfaces:

type TestInterface interface {
	sayHello() string
	getName() string
}

func greet(t TestInterface) string {
	return t.sayHello()
}

func interfaces() {
	println("\nInterfaces in Go can only have either functions or interfaces inside them.")
	println("So to mimic that behaviour we use getters and setters")

	println("\nUsing struct s1's object........")
	objectS1 := s1{
		name: "LightningMcAlan",
	}
	println(greet(objectS1))

	println("\nUsing struct s2's object........")
	objectS2 := s2{}
	println(greet(objectS2))
}

// Now we will define a struct to implement the Interface
// [Interfaces basically are a minimal definition for classes/structs to follow]
// For example: Dog, Cat etc.... Classes will implement interface Animal
// [No need for each of them to manually extend an Animal class]

type s1 struct {
	name string
}

func (s s1) sayHello() string {
	return "Hello! " + s.name
}

func (s s1) getName() string { // Assume it as a Getter function
	return s.name
}

type s2 struct{}

func (s s2) sayHello() string {
	return "My name is Maximus Decimus Meridius"
}

func (s s2) getName() string {
	return ""
}
