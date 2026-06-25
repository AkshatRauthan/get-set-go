package main

import (
	"slices"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"unit-converter/distance"
	"unit-converter/temperature"
	"unit-converter/weight"
)

// ---- Converter Interface ----
// Defined here in main — the consumer — not in a shared package.
// This is idiomatic Go: interfaces belong to the package that uses them,
// not the package that implements them. Each converter package independently
// satisfies this interface without importing it — implicit interface satisfaction.
// Coming from TS: this is structural typing, same as TypeScript's duck typing.

type Converter interface {
	Convert(value float64, from, to string) (float64, error)
}

// ---- Command Registry ----
// Maps the first CLI argument to a handler function.
// Adding a new command = one new entry here, one new handler function.
// No if-else chain, no giant switch in main.

type commandHandler func(args []string) error

var commands = map[string]commandHandler{
	"temperature": handleTemperature,
	"distance":    handleDistance,
	"weight":      handleWeight,
	"--help":      handleHelp,
	"-h":          handleHelp,
}

// ---- Entry Point ----

func main() {
	// No command given — show help
	if len(os.Args) < 2 {
		handleHelp(nil)
		os.Exit(0)
	}

	command := strings.ToLower(os.Args[1])

	handler, exists := commands[command]
	if !exists {
		fmt.Fprintf(os.Stderr, "Unknown command: '%s'\n\n", command)
		handleHelp(nil)
		os.Exit(1)
	}

	// Pass remaining args (everything after the command) to the handler
	err := handler(os.Args[2:])
	if err != nil {
		// Top level: extract ConversionError with errors.As — your instanceof pattern
		var ce *ConversionError
		if errors.As(err, &ce) {
			fmt.Fprintf(os.Stderr, "\nConversion failed:\n")
			fmt.Fprintf(os.Stderr, "  From : %s\n", ce.FromUnit)
			fmt.Fprintf(os.Stderr, "  To   : %s\n", ce.ToUnit)
			fmt.Fprintf(os.Stderr, "  Why  : %s\n\n", ce.Reason)
		} else {
			fmt.Fprintf(os.Stderr, "Error: %s\n\n", err.Error())
		}
		os.Exit(1)
	}
}

// ---- Conversion Handlers ----
// Each handler parses its own flags and delegates to the relevant converter.
// Flag format: --from <unit> --to <unit> --value <number>

func handleTemperature(args []string) error {
	value, from, to, err := parseConversionFlags(args)
	if err != nil {
		return fmt.Errorf("temperature: %w", err)
	}

	// Validate units before converting — gives a ConversionError with field info
	if !isValidUnit(from, temperature.SupportedUnits) {
		return newConversionError(from, to,
			fmt.Sprintf("'%s' is not a supported temperature unit. Use: %s",
				from, strings.Join(temperature.SupportedUnits, ", ")))
	}
	if !isValidUnit(to, temperature.SupportedUnits) {
		return newConversionError(from, to,
			fmt.Sprintf("'%s' is not a supported temperature unit. Use: %s",
				to, strings.Join(temperature.SupportedUnits, ", ")))
	}

	var c Converter = temperature.Converter{}
	result, err := c.Convert(value, from, to)
	if err != nil {
		return newConversionError(from, to, err.Error())
	}

	printResult(value, from, result, to)
	return nil
}

func handleDistance(args []string) error {
	value, from, to, err := parseConversionFlags(args)
	if err != nil {
		return fmt.Errorf("distance: %w", err)
	}

	if !isValidUnit(from, distance.SupportedUnits) {
		return newConversionError(from, to,
			fmt.Sprintf("'%s' is not a supported distance unit. Use: %s",
				from, strings.Join(distance.SupportedUnits, ", ")))
	}
	if !isValidUnit(to, distance.SupportedUnits) {
		return newConversionError(from, to,
			fmt.Sprintf("'%s' is not a supported distance unit. Use: %s",
				to, strings.Join(distance.SupportedUnits, ", ")))
	}

	var c Converter = distance.Converter{}
	result, err := c.Convert(value, from, to)
	if err != nil {
		return newConversionError(from, to, err.Error())
	}

	printResult(value, from, result, to)
	return nil
}

func handleWeight(args []string) error {
	value, from, to, err := parseConversionFlags(args)
	if err != nil {
		return fmt.Errorf("weight: %w", err)
	}

	if !isValidUnit(from, weight.SupportedUnits) {
		return newConversionError(from, to,
			fmt.Sprintf("'%s' is not a supported weight unit. Use: %s",
				from, strings.Join(weight.SupportedUnits, ", ")))
	}
	if !isValidUnit(to, weight.SupportedUnits) {
		return newConversionError(from, to,
			fmt.Sprintf("'%s' is not a supported weight unit. Use: %s",
				to, strings.Join(weight.SupportedUnits, ", ")))
	}

	var c Converter = weight.Converter{}
	result, err := c.Convert(value, from, to)
	if err != nil {
		return newConversionError(from, to, err.Error())
	}

	printResult(value, from, result, to)
	return nil
}

// ---- Flag Parser ----
// Parses --value <n> --from <unit> --to <unit> from args slice.
// Returns ErrMissingArgs if any required flag is absent.
// Returns ErrInvalidValue if --value cannot be parsed as float64.

func parseConversionFlags(args []string) (value float64, from string, to string, err error) {
	flags := parseFlags(args)

	rawValue, ok := flags["--value"]
	if !ok {
		return 0, "", "", fmt.Errorf("parseFlags: %w: --value is required", ErrMissingArgs)
	}

	value, err = strconv.ParseFloat(rawValue, 64)
	if err != nil {
		return 0, "", "", fmt.Errorf("parseFlags: %w: --value must be a number, got '%s'",
			ErrInvalidValue, rawValue)
	}

	from, ok = flags["--from"]
	if !ok {
		return 0, "", "", fmt.Errorf("parseFlags: %w: --from is required", ErrMissingArgs)
	}

	to, ok = flags["--to"]
	if !ok {
		return 0, "", "", fmt.Errorf("parseFlags: %w: --to is required", ErrMissingArgs)
	}

	return value, strings.ToLower(from), strings.ToLower(to), nil
}

// parseFlags converts a flat []string of --key value pairs into a map.
// e.g. ["--from", "celsius", "--to", "fahrenheit"] => {"--from": "celsius", "--to": "fahrenheit"}
func parseFlags(args []string) map[string]string {
	flags := make(map[string]string)
	for i := 0; i < len(args)-1; i++ {
		if strings.HasPrefix(args[i], "--") {
			flags[args[i]] = args[i+1]
			i++ // skip the value token
		}
	}
	return flags
}

// ---- Helpers ----

func isValidUnit(unit string, supported []string) bool {
	return slices.Contains(supported, unit)
}

func printResult(value float64, from string, result float64, to string) {
	fmt.Printf("\n  %.4f %s  =>  %.4f %s\n\n", value, from, result, to)
}

// ---- Help ----

func handleHelp(_ []string) error {
	fmt.Print(`
Unit Converter CLI
==================

USAGE:
  go run . <command> --value <number> --from <unit> --to <unit>

COMMANDS:
  temperature     Convert between temperature units
  distance        Convert between distance units
  weight          Convert between weight units
  --help, -h      Show this help message

FLAGS (required for all conversion commands):
  --value <number>    The numeric value to convert
  --from  <unit>      The unit to convert from
  --to    <unit>      The unit to convert to

TEMPERATURE UNITS:
  celsius       Degrees Celsius    (e.g. 100 celsius)
  fahrenheit    Degrees Fahrenheit (e.g. 212 fahrenheit)
  kelvin        Kelvin             (e.g. 373.15 kelvin)

DISTANCE UNITS:
  km            Kilometers         (e.g. 1 km)
  miles         Miles              (e.g. 0.621 miles)
  meters        Meters             (e.g. 1000 meters)

WEIGHT UNITS:
  kg            Kilograms          (e.g. 1 kg)
  pounds        Pounds             (e.g. 2.205 pounds)
  grams         Grams              (e.g. 1000 grams)

EXAMPLES:
  go run . temperature --value 100 --from celsius --to fahrenheit
  go run . temperature --value 98.6 --from fahrenheit --to celsius
  go run . distance    --value 5    --from km      --to miles
  go run . distance    --value 1    --from miles   --to meters
  go run . weight      --value 70   --from kg      --to pounds
  go run . weight      --value 500  --from grams   --to kg
`)
	return nil
}
