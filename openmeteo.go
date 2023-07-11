package cwf

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func MakeWeekUrl(city *City, mode string) ([]string, []string, error) {
	Url := "https://api.open-meteo.com/v1/forecast?latitude=" + city.Lat + "&longitude=" + city.Lng + "&daily=weathercode&timezone=Asia%2FTokyo"
	return makeRecest(Url, mode)
}
func MakedayUrl(city *City, mode string) ([]string, []string, error) {
	Url := "https://api.open-meteo.com/v1/forecast?latitude=" + city.Lat + "&longitude=" + city.Lng + "&hourly=weathercode&timezone=Asia%2FTokyo&forecast_days=1"
	return makeRecest(Url, mode)
}

func makeRecest(Url string, mode string) ([]string, []string, error) {
	request, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		fmt.Println("リクエストの作成に失敗しました:", err)
		return nil, nil, err
	}
	return makeClient(request, mode)
}

func makeClient(Request *http.Request, mode string) ([]string, []string, error) {
	// HTTPクライアントの作成
	client := &http.Client{}

	// リクエストの送信
	response, err := client.Do(Request)
	if err != nil {
		fmt.Println("リクエストの送信に失敗しました:", err)
		return nil, nil, err
	}
	defer response.Body.Close()
	if mode == "w" {
		return readWeekRespons(response)
	} else {
		return readDayRespons(response)
	}
}
func readWeekRespons(response *http.Response) ([]string, []string, error) {
	// レスポンスの読み取り
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("レスポンスの読み取りに失敗しました:", err)
		return nil, nil, err
	}
	// レスポンスを整形
	body_list := makeWeekBodyList(string(body))
	date := strings.Split(body_list[13], ",")
	weathercode := strings.Split(body_list[15], ",")
	return date, weathercode, err
}

func makeWeekBodyList(str_body string) []string {
	deleteChar := []string{"}", "[", "]", `"`}
	for _, char := range deleteChar {
		str_body = strings.Replace(str_body, char, "", -1)
	}
	str_body = strings.Replace(str_body, ",w", ":w", -1)
	return strings.Split(str_body, ":")
}

func readDayRespons(response *http.Response) ([]string, []string, error) {
	// レスポンスの読み取り
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("レスポンスの読み取りに失敗しました:", err)
		return nil, nil, err
	}
	// レスポンスを整形
	body_list := makeDayBodyList(string(body))
	// for i, j := range body_list {
	// 	fmt.Println(i, j)
	// }
	date := body_list[13:37]
	weathercode := strings.Split(body_list[38], ",")
	return date, weathercode, err
}

func makeDayBodyList(str_body string) []string {
	deleteChar := []string{"}", "[", "]", `"`, "00,"}
	for _, char := range deleteChar {
		str_body = strings.Replace(str_body, char, "", -1)
	}
	str_body = strings.Replace(str_body, ",w", ":w", -1)
	return strings.Split(str_body, ":")
}
