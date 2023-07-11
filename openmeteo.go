package cwf

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ResultWeek struct {
	WeekInf *Date `json:"daily"`
}
type ResultDay struct {
	DayInf *Date `json:"hourly"`
}
type Date struct {
	Time        []string `json:"time"`
	Weathercode []int    `json:"weathercode"`
}

func MakeUrl(city *City, mode string) ([]string, []int, error) {
	var Url string
	switch {
	case mode == "w":
		Url = "https://api.open-meteo.com/v1/forecast?latitude=" + city.Lat + "&longitude=" + city.Lng + "&daily=weathercode&timezone=Asia%2FTokyo"
	case mode == "d":
		Url = "https://api.open-meteo.com/v1/forecast?latitude=" + city.Lat + "&longitude=" + city.Lng + "&hourly=weathercode&timezone=Asia%2FTokyo&forecast_days=1"
	}
	return makeRecest(Url, mode)
}

func makeRecest(Url string, mode string) ([]string, []int, error) {
	request, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		fmt.Println("リクエストの作成に失敗しました:", err)
		return nil, nil, err
	}
	return makeClient(request, mode)
}

func makeClient(Request *http.Request, mode string) ([]string, []int, error) {
	// HTTPクライアントの作成
	client := &http.Client{}

	// リクエストの送信
	response, err := client.Do(Request)
	if err != nil {
		fmt.Println("リクエストの送信に失敗しました:", err)
		return nil, nil, err
	}
	defer response.Body.Close()

	return readRespons(response, mode)
}
func readRespons(response *http.Response, mode string) ([]string, []int, error) {
	// レスポンスの読み取り
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("レスポンスの読み取りに失敗しました:", err)
		return nil, nil, err
	}
	// レスポンスを整形
	return takeOutInf(body, mode)
}
func takeOutInf(body []uint8, mode string) ([]string, []int, error) {
	// レスポンスを整形
	var resultWeek ResultWeek
	var resultDay ResultDay
	var result *Date
	var err error

	switch {
	case mode == "w":
		err = json.Unmarshal(body, &resultWeek)
		result = resultWeek.WeekInf

	case mode == "d":
		err = json.Unmarshal(body, &resultDay)
		result = resultDay.DayInf
	}
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil, err
	}
	time := result.Time
	weathercode := result.Weathercode
	return time, weathercode, err
}
