package weather

//
// Types for Geocoding API

type GeocodingResult struct {
	City      string  `json:"name"`
	Region    string  `json:"admin1"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Elevation float32 `json:"elevation"`
}

type geocodingApiResponse struct {
	Results []GeocodingResult `json:"results"`
}

//
// Types for Current Weather

type CurrentWeatherResult struct {
	Time                string  `json:"time"`
	WeatherCode         uint8   `json:"weather_code"`
	WindSpeed           float32 `json:"wind_speed_10m"`
	Temperature         float32 `json:"temperature_2m"`
	RelativeHumidity    float32 `json:"relative_humidity_2m"`
	ApparentTemperature float32 `json:"apparent_temperature"`
}

type currentWeatherApiResponse struct {
	CurrentWeather CurrentWeatherResult `json:"current"`
}

//
// Types for Hourly Forecast

type HourlyWeatherResult struct {
	Timestamps               []string  `json:"time"`
	UvIndex                  []float32 `json:"uv_index"`
	Visibility               []float32 `json:"visibility"`
	Temperature              []float32 `json:"temperature_2m"`
	WindSpeed                []float32 `json:"wind_speed_10m"`
	PrecipitationProbability []float32 `json:"precipitation_probability"`
}

type hourlyWeatherApiResponse struct {
	HourlyWeather HourlyWeatherResult `json:"hourly"`
}

//
// Types for WeeklyForecast

type WeeklyWeatherResult struct {
	Date           []string  `json:"time"`
	MaxTemperature []float32 `json:"temperature_2m_max"`
	MinTemperature []float32 `json:"temperature_2m_min"`
	Precipitation  []float32 `json:"precipitation_sum"`
	Sunrise        []string  `json:"sunrise"`
	Sunset         []string  `json:"sunset"`
	MaxWindSpeed   []float32 `json:"wind_speed_10m_max"`
	MinWindSpeed   []float32 `json:"wind_speed_10m_min"`
	UvIndexMax     []float32 `json:"uv_index_max"`
}

type weeklyWeatherApiResponse struct {
	WeeklyWeather WeeklyWeatherResult `json:"daily"`
}

//
// Units

type Units string

const (
	Probability      Units = "%"
	RelativeHumidity Units = "%"
	Visibility       Units = "m"
	Temperature      Units = "°C"
	Precipitation    Units = "mm"
	WindSpeed        Units = "km/h"
)

//
// Weather Codes

type CodeKeys uint8

const (
	Clear            CodeKeys = 0
	MainlyClear      CodeKeys = 1
	PartlyCloudy     CodeKeys = 2
	Overcast         CodeKeys = 3
	Fog              CodeKeys = 45
	Fog2             CodeKeys = 48
	Drizzle1         CodeKeys = 51
	Drizzle2         CodeKeys = 53
	Drizzle3         CodeKeys = 55
	FreezingDrizzle1 CodeKeys = 56
	FreezingDrizzle2 CodeKeys = 57
	Rain1            CodeKeys = 61
	Rain2            CodeKeys = 63
	Rain3            CodeKeys = 65
	FreezingRain1    CodeKeys = 66
	FreezingRain2    CodeKeys = 67
	SnowFall1        CodeKeys = 71
	SnowFall2        CodeKeys = 73
	SnowFall3        CodeKeys = 75
	SnowGrains       CodeKeys = 77
	RainShowers1     CodeKeys = 80
	RainShowers2     CodeKeys = 81
	RainShowers3     CodeKeys = 82
	SnowShowers1     CodeKeys = 85
	SnowShowers2     CodeKeys = 86
	Thunderstorm1    CodeKeys = 95
	Thunderstorm2    CodeKeys = 96
	Thunderstorm3    CodeKeys = 99
)

var WeatherCodeMap = map[CodeKeys]string{
	Clear: "Clear Sky", MainlyClear: "Mostly Clear Sky", PartlyCloudy: "Partly Cloudy Sky", Overcast: "Fully Cloudy Sky", Fog: "Dense Fog",
	Fog2: "Freezing Rime Fog", Drizzle1: "Light Drizzle", Drizzle2: "Moderate Drizzle", Drizzle3: "Heavy Drizzle", FreezingDrizzle1: "Light Drizzle With Freezing Winds",
	FreezingDrizzle2: "Heavy Drizzle With Freezing Winds", Rain1: "Light Rain", Rain2: "Moderate Rain", Rain3: "Heavy Rain", FreezingRain1: "Light Rain With Freezing Winds",
	FreezingRain2: "Heavy Rain With Freezing Winds", SnowFall1: "Light Snowfall", SnowFall2: "Moderate Snowfall", SnowFall3: "Heavy Snowfall",
	SnowGrains: "Snow Grains", RainShowers1: "Light Distributed Rain Showers", RainShowers2: "Moderate Distributes Rain Showers", RainShowers3: "Heavy Distributed Rain Showers",
	SnowShowers1: "Light Distributed Snow Showers", SnowShowers2: "Heavy Distributed Snow Showers", Thunderstorm1: "Slight Thunderstorm", Thunderstorm2: "Heavy Thunderstorm",
	Thunderstorm3: "Thunderstorm With Hail",
}

//
// Errors

type ErrorResponse struct {
	Message string `json:"reason"`
	Error   bool   `json:"error"`
}
