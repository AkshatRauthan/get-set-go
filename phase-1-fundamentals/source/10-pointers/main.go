package main

func main() {
	println("\nPointers: ")
	pointers()
	println("NOTE: In Go there is no pointer arithematics unlike C/C++")

	println("\n\nReferences: ")
	println("NOTE: In go everything is pass by value...  Pass by reference do not exists")
	println("But due to internal working they seem to act like an indirect pass by reference.")

	println("\n\nIn case of maps:")
	referencedMaps()

	println("\n\nSlices: ")
	println("NOTE: Here same things happens but there is a slight diffrence...")
	println("We know that slices can be reallocated upon memory overflow during an append operation ")
	println("Now when this happens the slice that is inside the called function will be reallocated and its underlying pointer gets updated")
	println("Due to which further changes after reallocation will not be persistent by default")
	println("Also the size and cap are also passed by value so increase in them will also be discarted")
	println("Append Operation: Not persistent")
	println("Updations: Persistent strictly before first reallocation")
	println("NOTE: Inside called function we cant use methods like cap or len. To use them we need to create a new slice or pass a slice pointer")
	referencedSlices()

}
