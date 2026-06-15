package main

import "fmt"

type Animal struct {
	name  string
	sound string
}

func (a Animal) speak() {
	println(a.sound + "!!!!\n" + a.sound + "!!!!")
}

type Dog struct {
	Animal;
	breed string;
}

func (d Dog) speak() {
	println("Whoooof!!!!\nWhooooof!!!!")
}

func (d Dog) bark() {
	d.speak()
}

func compositionPattern() {
	println("\nComposition Pattern: ")
	println("They are used in GO for implementing inheritence along without using the classical tree based approach")

	println("The most basic type of these patterns is directly embedding structs inside structs...")

	myDog := Dog {
		Animal: Animal {
			name: "tommy",
			sound: "bark",
		},
		breed: "pug",
	}

	fmt.Println("\n", myDog) // {{tommy bark} pug}
	myDog.bark() 			 // both of the statements return Woof
	myDog.speak()

	println("\n\nTherefore, with this pattern we can actually perform function overriding in structs too, but only btw functions of diffrent structs\n")
}
