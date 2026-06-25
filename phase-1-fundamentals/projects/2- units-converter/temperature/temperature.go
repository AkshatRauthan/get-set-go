package temperature

import "fmt"

// ---- Supported Units ----
// Declared as constants so the compiler catches typos in unit strings
// across the conversion table below.

const (
	Celsius    = "celsius"
	Fahrenheit = "fahrenheit"
	Kelvin     = "kelvin"
)

// SupportedUnits is exported so main can display it in --help output.
var SupportedUnits = []string{Celsius, Fahrenheit, Kelvin}

// ---- Converter Struct ----
// Holds no state — all methods are value receivers since nothing mutates.
// The struct exists to satisfy the Converter interface defined in main.

type Converter struct{}

// Convert converts a temperature value from one unit to another.
// All conversions go through Celsius as an intermediate step — avoids
// writing N*(N-1) direct conversion formulas for N units.
// Returns an error wrapping ErrUnknownUnit for unrecognised unit strings.
func (c Converter) Convert(value float64, from, to string) (float64, error) {
	// Normalise to Celsius first
	celsius, err := toCelsius(value, from)
	if err != nil {
		return 0, fmt.Errorf("temperature.Convert: %w", err)
	}

	// Convert from Celsius to target unit
	result, err := fromCelsius(celsius, to)
	if err != nil {
		return 0, fmt.Errorf("temperature.Convert: %w", err)
	}

	return result, nil
}

// toCelsius converts any supported unit to Celsius.
// Unexported — internal step, not part of public API.
func toCelsius(value float64, from string) (float64, error) {
	switch from {
	case Celsius:
		return value, nil
	case Fahrenheit:
		// °C = (°F − 32) × 5/9
		return (value - 32) * 5 / 9, nil
	case Kelvin:
		// °C = K − 273.15
		return value - 273.15, nil
	default:
		return 0, fmt.Errorf("toCelsius: unknown unit '%s'", from)
	}
}

// fromCelsius converts a Celsius value to any supported unit.
// Unexported — internal step, not part of public API.
func fromCelsius(celsius float64, to string) (float64, error) {
	switch to {
	case Celsius:
		return celsius, nil
	case Fahrenheit:
		// °F = (°C × 9/5) + 32
		return (celsius*9/5) + 32, nil
	case Kelvin:
		// K = °C + 273.15
		return celsius + 273.15, nil
	default:
		return 0, fmt.Errorf("fromCelsius: unknown unit '%s'", to)
	}
}
