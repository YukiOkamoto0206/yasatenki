package weather

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type Weather struct {
	ReportDatetime string    `json:"reportDatetime"`
	TimeSeries     []OneTime `json:"timeSeries"`
}

type OneTime struct {
	TimeDefines []string `json:"timeDefines"`
	Areas       []Area   `json:"areas"`
}

type Area struct {
	AreaName    string   `json:"areaame"`
	Weathers    []string `json:"weathers"`
	Temperature []string `json:"temps"`
}

//func main() {
//	result := GetWeather()
//	fmt.Println(result)
//}

func GetWeather() string {
	url := "https://www.jma.go.jp/bosai/forecast/data/forecast/340000.json"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Get Http Error:", err)
	}
	// レスポンスボディを読み込む
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("IO Read Error:", err)
	}
	// 読み込み終わったらレスポンスボディを閉じる
	defer response.Body.Close()
	info := make([]*Weather, 0)
	err = json.Unmarshal(body, &info)
	if err != nil {
		log.Fatal(err)
	}

	weather := info[0].TimeSeries[0].Areas[0].Weathers[0]

	if strings.ContainsAny(weather, "雨") {
		return "今日は雨です！\n傘を忘れずに持っていきましょう！！"
	}

	return ""
}
