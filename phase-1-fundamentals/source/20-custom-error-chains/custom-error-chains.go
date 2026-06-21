package main

import (
	"errors"
	"fmt"
)

/*
	20. Custom Error Chains:

	In TS you defined a CustomError class and used instanceof to detect it:
	  try {
	    throw new ValidationError("email", "invalid format")
	  } catch (err) {
	    if (err instanceof ValidationError) { handle specific case }
	    else { generic handling }
	  }

	Go's equivalent:
	1. Define a struct with an Error() method (satisfies the `error` interface)
	2. Return it from the innermost function
	3. Forward it through intermediate layers using fmt.Errorf("...: %w", err)
	4. At the top level, use errors.As() to detect and extract it — this is instanceof

	KEY DIFFERENCE from TS:
	In TS, throw propagates the actual error OBJECT automatically through every function.
	In Go, each intermediate function must explicitly forward with %w.
	If any middle layer uses %v instead of %w, errors.As() silently returns false
	even though the same error MESSAGE string still prints.
	No compile error. Just broken instanceof-equivalent behaviour at runtime.
*/

// ValidationError is our custom error type — equivalent to your TS CustomError class
// Exported because in real code, callers in other packages need to errors.As() against it
type ValidationError struct {
	Field   string // which field failed (e.g. "email", "amount", "userID")
	Message string // what went wrong
}

// Error() satisfies the built-in `error` interface — required for any custom error type
// This is the single method the error interface requires
func (e *ValidationError) Error() string {
	return fmt.Sprintf("ValidationError: field='%s' message='%s'", e.Field, e.Message)
}

// ---- 3-layer call chain ----

// Layer 1 (innermost): validateAge — creates and returns the ValidationError
// This is where the error originates — deepest in the call stack
func validateAge(age int) error {
	if age < 0 {
		// Return the concrete custom error type — wrapped in error interface
		// Note: returning nil directly here (not a typed nil) to avoid nil interface trap
		return &ValidationError{Field: "age", Message: fmt.Sprintf("must be non-negative, got %d", age)}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: fmt.Sprintf("unrealistic value: %d", age)}
	}
	return nil
}

// Layer 2 (middle): validateUser — calls validateAge, wraps any error it receives
// This layer adds its own context but preserves the original ValidationError via %w
func validateUser(name string, age int) error {
	if name == "" {
		return &ValidationError{Field: "name", Message: "cannot be empty"}
	}

	err := validateAge(age)
	if err != nil {
		// %w preserves the ValidationError in the chain for errors.As() at the top
		// This is the Go equivalent of: catch (err) { throw err } — re-throwing with context
		return fmt.Errorf("validateUser (name=%s): %w", name, err)
	}

	return nil
}

// Layer 3 (outermost): createUser — calls validateUser, wraps any error it receives
// Another layer of context added, but the ValidationError chain is still intact via %w
func createUser(name string, age int) error {
	err := validateUser(name, age)
	if err != nil {
		return fmt.Errorf("createUser: %w", err)
	}

	fmt.Printf("User created successfully: name=%s, age=%d\n", name, age)
	return nil
}

func CustomErrorChainDemo() {
	println("\n\n01. Custom Error Chain — errors.As() as instanceof:")

	// Case 1: valid user — no error
	println("\nCase 1: valid input (should succeed):")
	err := createUser("akshat", 21)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Case 2: invalid age — ValidationError created deep in validateAge,
	// wrapped through validateUser and createUser, extracted at the top with errors.As()
	println("\nCase 2: negative age (ValidationError originates in validateAge):")
	err = createUser("akshat", -5)
	if err != nil {
		fmt.Println("Full error chain:", err)

		// errors.As() is the Go equivalent of `err instanceof ValidationError`
		// It unwraps the chain looking for a *ValidationError anywhere inside
		var ve *ValidationError
		if errors.As(err, &ve) {
			// ve is now populated with the actual ValidationError fields
			// equivalent to: catch(err) { if(err instanceof ValidationError) { err.field, err.message } }
			fmt.Println("Extracted via errors.As():")
			fmt.Println("  Field:  ", ve.Field)
			fmt.Println("  Message:", ve.Message)
		} else {
			fmt.Println("Not a ValidationError")
		}
	}

	// Case 3: empty name — ValidationError created in validateUser itself
	println("\nCase 3: empty name (ValidationError originates in validateUser):")
	err = createUser("", 21)
	if err != nil {
		fmt.Println("Full error chain:", err)
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Println("Extracted via errors.As():")
			fmt.Println("  Field:  ", ve.Field)
			fmt.Println("  Message:", ve.Message)
		}
	}

	// Case 4: unrealistic age — different ValidationError message, same type detection
	println("\nCase 4: age=999 (different ValidationError message, same type):")
	err = createUser("alan", 999)
	if err != nil {
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Printf("errors.As() correctly detected ValidationError regardless of message: field=%s\n", ve.Field)
		}
	}
}

/*
	The Broken Version — %v Kills errors.As():

	This is the exact failure mode your TS experience won't warn you about.

	In TS, `throw err` re-throws the actual error OBJECT through the chain.
	instanceof works because the object reference travels unchanged.

	In Go, fmt.Errorf with %v converts the error to a plain string and creates a NEW error.
	The original *ValidationError is discarded — only its message string survives.
	errors.As() can no longer find it because the type information is gone.
	The error string STILL PRINTS correctly — so it looks like it's working. But it isn't.
*/

// validateAgeBroken — same as validateAge but middle layer will use %v incorrectly
func validateAgeBroken(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: fmt.Sprintf("must be non-negative, got %d", age)}
	}
	return nil
}

// validateUserBroken — uses %v instead of %w when forwarding the error
// This BREAKS the chain: ValidationError is converted to a string here
func validateUserBroken(name string, age int) error {
	err := validateAgeBroken(age)
	if err != nil {
		// WRONG: %v discards the ValidationError and creates a new plain string error
		// errors.As() at the top will return false — the type info is gone
		return fmt.Errorf("validateUser (name=%s): %v", name, err)
	}
	return nil
}

func createUserBroken(name string, age int) error {
	err := validateUserBroken(name, age)
	if err != nil {
		return fmt.Errorf("createUser: %w", err) // %w here doesn't help — chain already broken above
	}
	return nil
}

func CustomErrorChainBrokenDemo() {
	println("\n\n02. Broken Custom Error Chain — %v Kills errors.As():")

	err := createUserBroken("akshat", -5)
	if err != nil {
		// Error string still looks correct — this is what makes it a silent bug
		fmt.Println("Error string (looks fine):", err)

		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Println("errors.As() found ValidationError — should not happen")
		} else {
			// errors.As() returns false — the ValidationError was lost at %v in validateUserBroken
			fmt.Println("errors.As() returned false — ValidationError lost at the %v wrap")
			fmt.Println("The error MESSAGE is intact but the TYPE information is gone")
			fmt.Println("This is the silent failure mode: error prints correctly but instanceof-equivalent fails")
		}
	}

	fmt.Println("\nLesson: %w preserves the error chain for errors.As() and errors.Is()")
	fmt.Println("        %v converts to string — chain broken, type info lost, no compile warning")
}
