package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const apiKey = "4d8fb5b93d4af21d66a2948710284366"
const apiURL = "https://api.openweathermap.org/data/2.5/weather"

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure float64 `json:"pressure"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Clouds struct {
		All float64 `json:"all"`
	} `json:"clouds"`
}

func main() {
	// Prompt user to enter city name
	fmt.Print("Enter city name: ")
	var city string
	fmt.Scanln(&city)

	fmt.Println()

	// Create request URL
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", apiURL, city, apiKey)

	// Make API request
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Parse JSON response
	weather := weatherData{}
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatal(err)
	}

	// Encode weather data into JSON format
	jsonData, err := json.Marshal(weather)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current weather in %s:\n", weather.Name)
	fmt.Println()
	fmt.Printf("  Temperature: %.2f°C\n", weather.Main.Temp)
	fmt.Printf("  Pressure: %.2f hPa\n", weather.Main.Pressure)
	fmt.Printf("  Humidity: %.0f%%\n", weather.Main.Humidity)
	fmt.Printf("  Description: %s\n", weather.Weather[0].Description)
	fmt.Printf("  Wind Speed: %.1f m/s\n", weather.Wind.Speed)
	fmt.Printf("  Cloudiness: %.0f%%\n", weather.Clouds.All)
	fmt.Println()
	fmt.Println()
	fmt.Printf("JSON Format:")
	fmt.Println()
	fmt.Println()

	// Print weather information in JSON format
	fmt.Println(string(jsonData))
}
