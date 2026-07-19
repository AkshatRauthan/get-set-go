package main

import (
	"flag"
	"fmt"
	"os"

	"cli-weather-dashboard/weather"
)

func usage() {
	fmt.Println("Usage: weather-cli <command> <city>")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  current   Show current weather for a city")
	fmt.Println("  hourly    Show hourly forecast for a city")
	fmt.Println("  weekly    Show 7-day forecast for a city")
	fmt.Println()
	fmt.Println("Example: weather-cli current \"New York\"")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		usage()
		os.Exit(1)
	}

	command := args[0]
	city := args[1]

	var err error
	switch command {
	case "current":
		err = weather.GetCurrentWeather(city)
	case "hourly":
		err = weather.GetHourlyWeatherForecast(city)
	case "weekly":
		err = weather.GetWeeklyWeatherForecast(city)
	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		usage()
		os.Exit(1)
	}

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
