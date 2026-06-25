package distance

import "fmt"

// ---- Supported Units ----

const (
	Kilometers = "km"
	Miles      = "miles"
	Meters     = "meters"
)

// SupportedUnits is exported so main can display it in --help output.
var SupportedUnits = []string{Kilometers, Miles, Meters}

// ---- Conversion Factors to Meters ----
// All conversions go through meters as the base unit.
// Adding a new unit only requires adding one entry here — no new formulas.

var toMeters = map[string]float64{
	Kilometers: 1000.0,
	Miles:      1609.344,
	Meters:     1.0,
}

// ---- Converter Struct ----

type Converter struct{}

// Convert converts a distance value from one unit to another via meters.
// Returns an error for any unrecognised unit string.
func (c Converter) Convert(value float64, from, to string) (float64, error) {
	fromFactor, ok := toMeters[from]
	if !ok {
		return 0, fmt.Errorf("distance.Convert: unknown unit '%s'", from)
	}

	toFactor, ok := toMeters[to]
	if !ok {
		return 0, fmt.Errorf("distance.Convert: unknown unit '%s'", to)
	}

	// Convert: value → meters → target unit
	meters := value * fromFactor
	result := meters / toFactor

	return result, nil
}
