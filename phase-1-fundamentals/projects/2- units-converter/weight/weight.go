package weight

import "fmt"

// ---- Supported Units ----

const (
	Kilograms = "kg"
	Pounds    = "pounds"
	Grams     = "grams"
)

// SupportedUnits is exported so main can display it in --help output.
var SupportedUnits = []string{Kilograms, Pounds, Grams}

// ---- Conversion Factors to Grams ----
// All conversions go through grams as the base unit.
// Same map-based pattern as distance — easy to extend.

var toGrams = map[string]float64{
	Kilograms: 1000.0,
	Pounds:    453.592,
	Grams:     1.0,
}

// ---- Converter Struct ----

type Converter struct{}

// Convert converts a weight value from one unit to another via grams.
// Returns an error for any unrecognised unit string.
func (c Converter) Convert(value float64, from, to string) (float64, error) {
	fromFactor, ok := toGrams[from]
	if !ok {
		return 0, fmt.Errorf("weight.Convert: unknown unit '%s'", from)
	}

	toFactor, ok := toGrams[to]
	if !ok {
		return 0, fmt.Errorf("weight.Convert: unknown unit '%s'", to)
	}

	// Convert: value → grams → target unit
	grams := value * fromFactor
	result := grams / toFactor

	return result, nil
}
