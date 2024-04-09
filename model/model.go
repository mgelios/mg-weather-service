package model

import "time"

type GeoCoding struct {
	Name       string            `json:"name"`
	Lat        float64           `json:"lat"`
	Lon        float64           `json:"lon"`
	Country    string            `json:"country"`
	LocalNames map[string]string `json:"local_names"`
	State      string            `json:"state"`
	Created    time.Time         `json:"created"`
	QueredCity string            `json:"quered_city"`
}

type OneCallWeather struct {
	Lat            float64        `json:"lat"`
	Lon            float64        `json:"lon"`
	Timezone       string         `json:"timezone"`
	TimezoneOffset int32          `json:"timezone_offset"`
	CurrentWeather CurrentWeather `json:"current,omitempty"`
	Daily          []DailyWeather `json:"daily,omitempty"`
	Created        time.Time      `json:"created"`
}

type CurrentWeather struct {
	Daytime    int                  `json:"dt"`
	Sunrise    int                  `json:"sunrise"`
	Sunset     int                  `json:"sunset"`
	Temp       float64              `json:"temp"`
	FeelsLike  float64              `json:"feels_like"`
	Pressure   int64                `json:"pressure"`
	Humidity   int64                `json:"humidity"`
	DewPoint   float64              `json:"dew_point"`
	UVI        float64              `json:"uvi"`
	Clouds     int64                `json:"clouds"`
	Visibility int64                `json:"visibility"`
	WindSpeed  float64              `json:"wind_speed"`
	WindDeg    int64                `json:"wind_deg"`
	Weather    []WeatherDescription `json:"weather"`
}

type WeatherDescription struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type DailyWeather struct {
	Daytime    int                  `json:"dt"`
	Sunrise    int                  `json:"sunrise"`
	Sunset     int                  `json:"sunset"`
	Moonrise   int                  `json:"moonrise"`
	Moonset    int                  `json:"moonset"`
	Temp       DailyTemp            `json:"temp"`
	FeelsLike  DailyFeelsLike       `json:"feels_like"`
	Pressure   int64                `json:"pressure"`
	Humidity   int64                `json:"humidity"`
	DewPoint   float64              `json:"dew_point"`
	UVI        float64              `json:"uvi"`
	Clouds     int64                `json:"clouds"`
	Pop        float64              `json:"pop"`
	Visibility int64                `json:"visibility"`
	WindSpeed  float64              `json:"wind_speed"`
	WindDeg    int64                `json:"wind_deg"`
	WindGust   float64              `json:"wind_gust"`
	Weather    []WeatherDescription `json:"weather"`
}

type DailyTemp struct {
	Day     float64 `json:"day"`
	Night   float64 `json:"night"`
	Evening float64 `json:"eve"`
	Morning float64 `json:"morn"`
	Min     float64 `json:"min"`
	Max     float64 `json:"max"`
}

type DailyFeelsLike struct {
	Day     float64 `json:"day"`
	Night   float64 `json:"night"`
	Evening float64 `json:"eve"`
	Morning float64 `json:"morn"`
}
