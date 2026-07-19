package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// getRequestHelper wraps the logic to make GET request with TimeOut to the address url with queryParams as request query.
func getRequestHelper(apiUrl string, queryParams url.Values) ([]byte, error) {
	if apiUrl == "" {
		return nil, fmt.Errorf("getRequest: url cannot be empty")
	}
	client := http.Client{Timeout: 3 * time.Second}

	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("getRequest-1: %w", err)
	}

	req.URL.RawQuery = queryParams.Encode()

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("getRequest-2: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var body, _ = io.ReadAll(res.Body)
		var errorBody ErrorResponse
		if err = json.Unmarshal(body, &errorBody); err != nil {
			return nil, fmt.Errorf("getRequest-3: unexpected status code: %d", res.StatusCode)
		}
		return nil, fmt.Errorf("getRequest-4: %s", errorBody.Message)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// fetchCoordinates take a city as an input and returns its coordinates along with other metadata.
func fetchCoordinates(city string) (GeocodingResult, error) {
	apiUrl := "https://geocoding-api.open-meteo.com/v1/search"
	urlParams := make(map[string][]string)

	urlParams["name"] = []string{city}
	urlParams["count"] = []string{"1"}
	urlParams["language"] = []string{"en"}
	urlParams["format"] = []string{"json"}

	resBody, err := getRequestHelper(apiUrl, urlParams)
	if err != nil {
		return GeocodingResult{}, fmt.Errorf("fetchCoordinates-1: %w", err)
	}

	var apiResponse geocodingApiResponse
	err = json.Unmarshal(resBody, &apiResponse)
	if err != nil {
		return GeocodingResult{}, fmt.Errorf("fetchCoordinates-2: %w", err)
	}

	if len(apiResponse.Results) == 0 {
		return GeocodingResult{}, fmt.Errorf("fetchCoordinates-3: Invalid city name entered")
	}

	return apiResponse.Results[0], nil
}

// fetchWeeklyForecast take coordinates of a place and return its weekly weather forecast
func fetchWeeklyForecast(longitude float64, latitude float64) (WeeklyWeatherResult, error) {
	apiUrl := "https://api.open-meteo.com/v1/forecast"

	urlParams := make(map[string][]string)
	urlParams["latitude"] = []string{strconv.FormatFloat(latitude, 'f', -4, 32)}
	urlParams["longitude"] = []string{strconv.FormatFloat(longitude, 'f', -4, 32)}
	urlParams["daily"] = []string{
		"temperature_2m_max", "temperature_2m_min", "precipitation_sum", "sunrise", "sunset",
		"wind_speed_10m_max", "wind_speed_10m_min", "uv_index_max",
	}

	resBody, err := getRequestHelper(apiUrl, urlParams)
	if err != nil {
		return WeeklyWeatherResult{}, fmt.Errorf("fetchWeeklyForecast-1: %w", err)
	}

	var apiRes weeklyWeatherApiResponse
	err = json.Unmarshal(resBody, &apiRes)
	if err != nil {
		return WeeklyWeatherResult{}, fmt.Errorf("fetchWeeklyForecast-2: %w", err)
	}

	return apiRes.WeeklyWeather, nil
}

// fetchHourlyForecast take coordinates of a place and return its hourly forecast for today
func fetchHourlyForecast(longitude float64, latitude float64) (HourlyWeatherResult, error) {
	apiUrl := "https://api.open-meteo.com/v1/forecast"
	urlParams := make(map[string][]string)
	urlParams["latitude"] = []string{strconv.FormatFloat(latitude, 'f', -4, 32)}
	urlParams["longitude"] = []string{strconv.FormatFloat(longitude, 'f', -4, 32)}
	urlParams["hourly"] = []string{"temperature_2m", "precipitation_probability", "wind_speed_10m", "visibility", "uv_index"}

	resBody, err := getRequestHelper(apiUrl, urlParams)
	if err != nil {
		return HourlyWeatherResult{}, fmt.Errorf("fetchHourlyForecast-1: %w", err)
	}

	var apiRes hourlyWeatherApiResponse
	err = json.Unmarshal(resBody, &apiRes)
	if err != nil {
		return HourlyWeatherResult{}, fmt.Errorf("fetchHourlyForecast-2: %w", err)
	}

	return apiRes.HourlyWeather, nil
}

// fetchCurrentWeather take coordinates of a place and return its current weather information
func fetchCurrentWeather(longitude float64, latitude float64) (CurrentWeatherResult, error) {
	apiUrl := "https://api.open-meteo.com/v1/forecast"
	urlParams := make(map[string][]string)
	urlParams["latitude"] = []string{strconv.FormatFloat(latitude, 'f', -4, 32)}
	urlParams["longitude"] = []string{strconv.FormatFloat(longitude, 'f', -4, 32)}
	urlParams["current"] = []string{"temperature_2m", "relative_humidity_2m", "wind_speed_10m", "weather_code", "apparent_temperature"}

	resBody, err := getRequestHelper(apiUrl, urlParams)
	if err != nil {
		return CurrentWeatherResult{}, fmt.Errorf("fetchCurrentWeather-1: %w", err)
	}

	var apiRes currentWeatherApiResponse
	err = json.Unmarshal(resBody, &apiRes)
	if err != nil {
		return CurrentWeatherResult{}, fmt.Errorf("fetchCurrentWeather-2: %w", err)
	}
	return apiRes.CurrentWeather, nil
}
