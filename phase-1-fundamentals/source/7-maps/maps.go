package main

import "fmt"

// maps -> hashmaps
func maps() {
	// Initialising maps
	// m["int"] = "string"

	var m1 = make(map[int]string)
	m2 := make(map[string]int)
	m3 := map[string]int{"zero": 0, "ten": 10}

	// Clearing complete map
	clear(m1)
	clear(m2)

	// Inserting elements
	m1[1] = "one"
	m1[2] = "two"
	m2["one"] = 1
	m2["two"] = 2
	m3["six"] = 6

	fmt.Println("\nCreated Map m1: ", m1)
	fmt.Println("Created Map m2: ", m2)

	// IMP: if key doesn't exists it returns zeroed value of the value datatype.
	fmt.Println("\nAccessing uninitialised key: ", m1[10])  // ""
	fmt.Println("Accessing uninitialised key: ", m2["ten"]) // 0

	// Deleting elements
	delete(m1, 1)
	delete(m2, "one")

	fmt.Println("\nMap m1 after deletion: ", m1, "")
	fmt.Println("Map m2 after deletion: ", m2, "\n")

	// Checking if key exists in Map
	k11 := 1
	k12 := 2

	// Not exists Else will run
	val, ok := m1[k11] // ok -> false
	if ok {
		fmt.Println("Key: ", k11, " Exists in Map m1")
		fmt.Println("Value of ", k11, " is ", val)
	} else {
		fmt.Println("Key: ", k11, " Does not exists in Map m1")
	}

	// Exists If will run
	val, ok = m1[k12] // ok -> true
	if ok {
		fmt.Println("Key: ", k12, " Exists in Map m1")
		fmt.Println("Value of ", k12, " is ", val)
	} else {
		fmt.Println("Key: ", k12, " Does not exists in Map m1")
	}

}
