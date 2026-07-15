package main

import (
	"errors"
	"fmt"
)

// ---- Sentinel Errors ----

var ErrUnknownUnit    = errors.New("unknown unit")
var ErrSameUnit       = errors.New("from and to units are the same")
var ErrInvalidValue   = errors.New("invalid numeric value")
var ErrMissingArgs    = errors.New("missing required arguments")

// ---- Custom Error Type ----
// ConversionError carries the specific units that failed — equivalent to
// your TS CustomError class with field properties.
// Callers extract this with errors.As() — same as your instanceof check.

type ConversionError struct {
	FromUnit string
	ToUnit   string
	Reason   string
}

func (e *ConversionError) Error() string {
	return fmt.Sprintf(
		"ConversionError: cannot convert '%s' to '%s' — %s",
		e.FromUnit, e.ToUnit, e.Reason,
	)
}

// newConversionError is an unexported helper — internal convenience constructor.
// Callers outside main use the returned error interface, not this type directly.
func newConversionError(from, to, reason string) error {
	return &ConversionError{
		FromUnit: from,
		ToUnit:   to,
		Reason:   reason,
	}
}
