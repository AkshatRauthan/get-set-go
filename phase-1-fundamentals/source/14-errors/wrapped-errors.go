package main

import (
	"errors"
	"fmt"
)

/*
	03. Wrapped Errors:
	Wrapping an error means attaching additional context to an existing error without losing the original error.
	Instead of just returning a raw error, we wrap it with a message that explains what we were doing when it failed.
	The caller can then use errors.Is() to check for a specific underlying error anywhere in the chain,
	or errors.As() to extract a specific error type — both work even through multiple layers of wrapping.
*/

var RecordNotFoundError = errors.New("Error: Record not found in database")
var UnauthorizedAccessError = errors.New("Error: Unauthorized access to resource")

// fetchRecord simulates a DB lookup — returns a wrapped error with context on failure
func fetchRecord(userID int, resourceID int) (string, error) {
	// simulated DB: only record 101 exists
	db := map[int]string{101: "Project Alpha Report"}

	record, ok := db[resourceID]
	if !ok {
		// %w verb wraps the sentinel error, preserving it for errors.Is() checks
		return "", fmt.Errorf("fetchRecord (userID=%d, resourceID=%d): %w", userID, resourceID, RecordNotFoundError)
	}

	return record, nil
}

// authorizeAndFetch checks access rights, then delegates to fetchRecord — wraps any error it gets back
func authorizeAndFetch(userID int, resourceID int) (string, error) {
	// simulated access control: only userID 1 has access
	allowedUsers := map[int]bool{1: true}

	if !allowedUsers[userID] {
		return "", fmt.Errorf("authorizeAndFetch (userID=%d): %w", userID, UnauthorizedAccessError)
	}

	record, err := fetchRecord(userID, resourceID)
	if err != nil {
		// wrapping again => adds another layer of context on top of fetchRecord's wrapped error
		return "", fmt.Errorf("authorizeAndFetch: access granted but fetch failed: %w", err)
	}

	return record, nil
}

func WrappedErrors() {
	println("\n\n03. Wrapped Errors:")

	// Case 1: valid user, valid record
	println("\nFetching record 101 as user 1 (should succeed):")
	record, err := authorizeAndFetch(1, 101)
	if err != nil {
		println(err.Error())
	} else {
		println("Fetched record:", record)
	}

	// Case 2: valid user, record does not exist => errors.Is() unwraps the chain to find RecordNotFoundError
	println("\nFetching record 999 as user 1 (record missing):")
	_, err = authorizeAndFetch(1, 999)
	if err != nil {
		println("Raw wrapped error =>", err.Error())
		if errors.Is(err, RecordNotFoundError) {
			println("Confirmed via errors.Is() => RecordNotFoundError is in the chain")
		}
		fmt.Println("\nPrinting an Wrapped Error Object:\n", err.Error())
	}

	// Case 3: unauthorized user => errors.Is() catches UnauthorizedAccessError
	println("\nFetching record 101 as user 42 (unauthorized):")
	_, err = authorizeAndFetch(42, 101)
	if err != nil {
		println("Raw wrapped error =>", err.Error())
		if errors.Is(err, UnauthorizedAccessError) {
			println("Confirmed via errors.Is() => UnauthorizedAccessError is in the chain")
		}
		fmt.Println("\nPrinting an Wrapped Error Object:\n", err.Error())
	}
}
