package main
import (
    _ "embed"
    "encoding/json"
    "fmt"
    "strconv"
)

//go:embed "cities.json"
var citiesJson []byte
type City struct {
    Country string `json:"country"`
    Name string `json:"name"`
    Lat string `json:"Lat"`
    Lng string `json:"Lng"`
    Latitude float64 `json:"-"`
    Longitude float64 `json:"-”`
}
func main() {
	cities := []*City{}
	err := json.Unmarshal(citiesJson, &cities) 
	if err != nil {
        fmt.Println(err.Error())
		return
	}
	for _, city := range cities {
		city.Latitude, _ = strconv.ParseFloat(city.Lat, 64)
		city.Longitude, _ = strconv.ParseFloat(city.Lng, 64)
		fmt.Printf("Lat %f, Lng %f\n", city.Latitude, city.Longitude)
	}
	fmt.Printf("read %d entries\n", len(cities))
	cityinfo := week("Tokyo", cities)
	fmt.Printf("%s\n", cityinfo.Name)
}
func week(cityname string, cities []*City) City{
	var cityinfo City
	for _,city := range(cities){
		if city.Name == cityname {
			fmt.Printf("%s, Lat: %f, Lng: %f\n",city.Name, city.Latitude, city.Longitude)
			cityinfo = *city
			break
		}
	}
	return cityinfo
}