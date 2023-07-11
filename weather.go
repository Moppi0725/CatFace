package cwf

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed "weather.json"
var weatherJson []byte

type Weather struct {
	Code string `json:"code"`
	Jp   string `json:"jp"`
}

func GetWeatherMap() (map[string]string, error) {
	weather := []*Weather{}
	err := json.Unmarshal(weatherJson, &weather)
	weatherMap := map[string]string{}
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	for _, i := range weather {
		weatherMap[i.Code] = i.Jp
	}
	return weatherMap, nil
}

func OutputWeather(weatherMap map[string]string, code int) (string, error) {
	if _, weatherBoolean := weatherMap[fmt.Sprint(code)]; weatherBoolean {
		return weatherMap[fmt.Sprint(code)], nil
	}
	return "", fmt.Errorf("%s: WeatherCode no found", fmt.Sprint(code))
}
