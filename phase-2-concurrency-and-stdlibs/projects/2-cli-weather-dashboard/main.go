package main

import (
	"cli-weather-dashboard/weather"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// If All API calls don't get completed within 5sec for any req we will get timeout error
	// Here only network calls (http req) are watching the context.
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

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
		err = weather.GetCurrentWeather(ctx, city)
	case "hourly":
		err = weather.GetHourlyWeatherForecast(ctx, city)
	case "weekly":
		err = weather.GetWeeklyWeatherForecast(ctx, city)
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
