package cwf

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed "cities.json"
var citiesJson []byte

type City struct {
	Country string `json:"country"`
	Name    string `json:"name"`
	Lat     string `json:"Lat"`
	Lng     string `json:"Lng"`
}

func GetCityInfo(cityname string) (*City, error) {
	cities := []*City{}
	err := json.Unmarshal(citiesJson, &cities)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return serchCity(cityname, cities)
}

func serchCity(cityname string, cities []*City) (*City, error) {
	for _, city := range cities {
		if city.Name == cityname {
			return city, nil
		}
	}

	return nil, fmt.Errorf("%s: city no found", cityname)
}
