package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
)

type Weather struct {
	Status WeatherStatus `json:"status"`
}

type WeatherStatus struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type WeatherInfo struct {
	Water       int    `json:"water"`
	WaterStatus string `json:"waterStatus"`
	Wind        int    `json:"wind"`
	WindStatus  string `json:"windStatus"`
}

var PORT = ":3000"

func main() {
	http.HandleFunc("/", templateHandler)

	http.ListenAndServe(PORT, nil)
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	// push(w, "/assets/css/style.scss")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to load template")
	}

	info := updateWeather()
	// fmt.Println(info)

	t.Execute(w, info)
}

func updateWeather() WeatherInfo {
	water := rand.Intn(99) + 1
	wind := rand.Intn(99) + 1

	weather := Weather{}

	weather.Status.Wind = wind
	weather.Status.Water = water

	updateFile(weather, "weathers.json")

	return getStatusInfo(water, wind)
}

func getStatusInfo(water int, wind int) WeatherInfo {
	weatherInfo := WeatherInfo{
		Water:       water,
		WaterStatus: getWaterStatus(water),
		Wind:        wind,
		WindStatus:  getWindStatus(wind),
	}

	return weatherInfo
}

func getWaterStatus(water int) string {
	if water <= 5 {
		return "Safe"
	} else if water >= 6 && water <= 8 {
		return "Warning"
	} else {
		return "Danger"
	}
}

func getWindStatus(wind int) string {
	if wind <= 6 {
		return "Safe"
	} else if wind >= 7 && wind <= 15 {
		return "Warning"
	} else {
		return "Danger"
	}
}

// func getWeatherJsonData() *Weather {
// 	jsonFile, err := os.Open("weathers.json")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer jsonFile.Close()

// 	weather, err := readFile(jsonFile)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// fmt.Println(weathers)

// 	return weather
// }

// func readFile(data *os.File) (*Weather, error) {
// 	rawData, err := ioutil.ReadAll(data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var weather *Weather
// 	err = json.Unmarshal(rawData, &weather)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return weather, nil
// }

func updateFile(data Weather, filename string) error {
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, value, 0644)
}
