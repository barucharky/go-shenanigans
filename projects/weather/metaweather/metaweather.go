// B''H

package metaweather

const LocationURL = "https://www.metaweather.com/api/location/"

type LocationSearchResult struct {
	Title        string
	LocationType string `json:"location_type"`
	Woeid        int
}

type WeatherResult struct {
	Title string
	Items []*FiveDay `json:"consolidated_weather"`
}

type FiveDay struct {
	Condition     string  `json:"weather_state_name"`
	WindDirection string  `json:"wind_direction_compass"`
	Date          string  `json:"applicable_date"`
	MinTemp       float64 `json:"min_temp"`
	MaxTemp       float64 `json:"max_temp"`
	Humidity      int
}

//!-
