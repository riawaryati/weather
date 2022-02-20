package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
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

var PORT = ":8080"

func main() {
	http.HandleFunc("/", templateHandler)

	http.ListenAndServe(PORT, nil)
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/assets/css/style.scss")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to load template")
	}

	info := updateWeather()

	t.Execute(w, info)
}

func push(w http.ResponseWriter, resource string) {
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push(resource, nil); err == nil {
			return
		}
	}
}

func updateWeather() WeatherInfo {
	weather := getWeatherJsonData()

	water := rand.Intn(99) + 1
	wind := rand.Intn(99) + 1

	weather.Status.Water = water
	weather.Status.Wind = wind

	updateFile(*weather, "weathers.json")

	return getStatusInfo(*weather)
}

func getStatusInfo(weater Weather) WeatherInfo {

	weatherInfo := WeatherInfo{
		Water:       weater.Status.Water,
		WaterStatus: getWaterStatus(weater.Status.Water),
		Wind:        weater.Status.Wind,
		WindStatus:  getWindStatus(weater.Status.Wind),
	}

	return weatherInfo
}

func getWaterStatus(water int) string {
	if water < 5 {
		return "Safe"
	} else if water >= 6 && water <= 8 {
		return "Warning"
	} else {
		return "Danger"
	}
}

func getWindStatus(wind int) string {
	if wind < 6 {
		return "Safe"
	} else if wind >= 7 && wind <= 15 {
		return "Warning"
	} else {
		return "Danger"
	}
}

func getWeatherJsonData() *Weather {
	jsonFile, err := os.Open("weathers.json")
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	weather, err := readFile(jsonFile)
	if err != nil {
		panic(err)
	}

	// fmt.Println(weathers)

	return weather
}

func readFile(data *os.File) (*Weather, error) {
	rawData, err := ioutil.ReadAll(data)
	if err != nil {
		return nil, err
	}

	var weather *Weather
	err = json.Unmarshal(rawData, &weather)
	if err != nil {
		return nil, err
	}

	return weather, nil
}

func updateFile(data Weather, filename string) error {
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, value, 0644)
}
