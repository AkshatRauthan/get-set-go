package main

import "errors"

/*
	02. Sentinel Errors:
	A sentinel error is a pre-declared, package-level error variable that indicates a specific, expected failure state.
	Instead of just returning a generic error string that the caller has to parse, we return a specific named error.
	This allows the caller to check exactly what went wrong using equality, rather than messy string matching.
*/

// to be defined exacttly like this: no shorthand declaration(only works inside func calls), no const keyword
var KeyDoNotExistsError = errors.New("Error: Requested key does not exist")
var KeyAlreadyExistsError = errors.New("Error: Requested key already exists")

func wrapper() func(m map[int]string, key int, value string, updateVal bool) (map[int]string, error) {

	// updateVal => true updating, flase => inserting
	return func(m map[int]string, key int, value string, updateVal bool) (map[int]string, error) {

		if updateVal {
			// for updation => key must exists
			_, ok := m[key]
			if !ok {
				return m, KeyDoNotExistsError // Throwing Error
			}
		} else {
			// for insertion => key must not exists
			_, ok := m[key]
			if ok {
				return m, KeyAlreadyExistsError // Throwing Error
			}
		}

		m[key] = value
		return m, nil
	}
}

func SentinelErrors() {
	println("\n\n02. Sentinel Errors:")

	m := make(map[int]string)
	insertValueInMap, updateValueInMap := wrapper(), wrapper()

	println("\nInserting 1:'one' into map:")
	m, err := insertValueInMap(m, 1, "one", false)
	if err != nil {
		println(err.Error())
	} else {
		println("Insertion Successfull...")
	}

	println("\nInserting 1:'two' into map:")
	m, err = insertValueInMap(m, 1, "two", false)
	if err != nil {
		println(err.Error())
	} else {
		println("Insertion Successfull...")
	}

	println("\nUpdating 1's value to 'two' into map:")
	m, err = updateValueInMap(m, 1, "two", true)
	if err != nil {
		println(err.Error())
	} else {
		println("Updation Successfull...")
	}

	println("\nUpdating 2's value to 'two' into map:")
	m, err = updateValueInMap(m, 2, "two", true)
	if err != nil {
		println(err.Error())
	} else {
		println("Updation Successfull...")
	}

}
