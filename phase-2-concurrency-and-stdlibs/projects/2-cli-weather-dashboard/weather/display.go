package weather

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func GetCurrentWeather(ctx context.Context, cityName string) error {
	cityDetails, err := fetchCoordinates(ctx, cityName)
	if err != nil {
		return fmt.Errorf("GetCurrentWeather-1: %v", err)
	}
	currentWeather, err := fetchCurrentWeather(ctx, cityDetails.Longitude, cityDetails.Latitude)
	if err != nil {
		return fmt.Errorf("GetCurrentWeather-2: %v", err)
	}
	displayCurrentWeather(cityDetails, currentWeather)
	return nil
}

func GetHourlyWeatherForecast(ctx context.Context, cityName string) error {
	cityDetails, err := fetchCoordinates(ctx, cityName)
	if err != nil {
		return fmt.Errorf("GetHourlyWeatherForecast-1: %v", err)
	}
	hourlyForecast, err := fetchHourlyForecast(ctx, cityDetails.Longitude, cityDetails.Latitude)
	if err != nil {
		return fmt.Errorf("GetHourlyWeatherForecast-2: %v", err)
	}
	displayHourlyForecast(cityDetails, hourlyForecast)
	return nil
}

func GetWeeklyWeatherForecast(ctx context.Context, cityName string) error {
	cityDetails, err := fetchCoordinates(ctx, cityName)
	if err != nil {
		return fmt.Errorf("GetWeeklyWeatherForecast-1: %v", err)
	}
	weeklyForecast, err := fetchWeeklyForecast(ctx, cityDetails.Longitude, cityDetails.Latitude)
	if err != nil {
		return fmt.Errorf("GetWeeklyWeatherForecast-2: %v", err)
	}
	displayWeeklyForecast(cityDetails, weeklyForecast)
	return nil
}

//
// Display helpers
//

func printHeader(city GeocodingResult, title string) {
	location := city.City
	if city.Region != "" {
		location += ", " + city.Region
	}
	location += ", " + city.Country

	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("%s — %s\n", title, location)
	fmt.Println(strings.Repeat("=", 60))
}

func displayCurrentWeather(city GeocodingResult, w CurrentWeatherResult) {
	printHeader(city, "Current Weather")

	condition := WeatherCodeMap[CodeKeys(w.WeatherCode)]
	if condition == "" {
		condition = "Unknown"
	}

	fmt.Printf("Condition:     %s\n", condition)
	fmt.Printf("Temperature:   %.1f%s (feels like %.1f%s)\n", w.Temperature, Temperature, w.ApparentTemperature, Temperature)
	fmt.Printf("Humidity:      %.0f%s\n", w.RelativeHumidity, RelativeHumidity)
	fmt.Printf("Wind Speed:    %.1f %s\n", w.WindSpeed, WindSpeed)
	fmt.Printf("Observed At:   %s\n", formatTimestamp(w.Time))
	fmt.Println(strings.Repeat("-", 50))
}

func displayHourlyForecast(city GeocodingResult, h HourlyWeatherResult) {
	printHeader(city, "Hourly Forecast")

	fmt.Printf("%-18s %10s %8s %10s %6s\n", "Time", "Temp", "Precip%", "Wind", "UV")
	fmt.Println(strings.Repeat("-", 60))

	for i := range h.Timestamps {
		fmt.Printf(
			"%-18s %9.1f%s %7.0f%s %8.1f%s %6.1f\n",
			formatTimestamp(h.Timestamps[i]),
			h.Temperature[i], Temperature,
			h.PrecipitationProbability[i], Probability,
			h.WindSpeed[i], WindSpeed,
			h.UvIndex[i],
		)
	}
	fmt.Println(strings.Repeat("-", 60))
}

func displayWeeklyForecast(city GeocodingResult, wk WeeklyWeatherResult) {
	printHeader(city, "7-Day Forecast")

	fmt.Printf("%-12s %10s %10s %10s %8s %8s %6s\n", "Date", "High", "Low", "Precip", "Sunrise", "Sunset", "UV")
	fmt.Println(strings.Repeat("-", 70))

	for i := range wk.Date {
		fmt.Printf(
			"%-12s %9.1f%s %9.1f%s %8.1f%s %8s %8s %6.1f\n",
			formatDate(wk.Date[i]),
			wk.MaxTemperature[i], Temperature,
			wk.MinTemperature[i], Temperature,
			wk.Precipitation[i], Precipitation,
			formatClockTime(wk.Sunrise[i]),
			formatClockTime(wk.Sunset[i]),
			wk.UvIndexMax[i],
		)
	}
	fmt.Println(strings.Repeat("-", 70))
}

//
// Time formatting helpers

// formatTimestamp turns "2026-07-19T14:00" into "Jul 19, 2:00 PM"
func formatTimestamp(raw string) string {
	t, err := time.Parse("2006-01-02T15:04", raw)
	if err != nil {
		return raw
	}
	return t.Format("Jan 2, 3:04 PM")
}

// formatDate turns "2026-07-19" into "Sun, Jul 19"
func formatDate(raw string) string {
	t, err := time.Parse("2006-01-02", raw)
	if err != nil {
		return raw
	}
	return t.Format("Mon, Jan 2")
}

// formatClockTime turns "2026-07-19T05:47" into "5:47 AM"
func formatClockTime(raw string) string {
	t, err := time.Parse("2006-01-02T15:04", raw)
	if err != nil {
		return raw
	}
	return t.Format("3:04 PM")
}
