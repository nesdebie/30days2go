package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)


type APIConfig struct {
	ApiKey string
}


type WeatherResponse struct {
	CityName string `json:"name"`
	Main     struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
	Timezone int `json:"timezone"`
}


type FinalResponse struct {
	Name      string  `json:"name"`
	Temp      float64 `json:"temp"`
	LocalTime string  `json:"local_time"`
}


func loadConfig() (APIConfig, error) {
	var config APIConfig

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return config, err
	}

	key, exists := os.LookupEnv("API_KEY")
	if !exists {
		log.Fatal("Missing API key in environment variables")
		return config, errors.New("missing API key")
	}

	config = APIConfig{ApiKey: key}
	return config, nil
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "Usage: curl localhost:8080/weather/<city_name>", http.StatusBadRequest)
		return
	}

	city := pathParts[2]

	data, err := getTemperature(city)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch city data: %v", err), http.StatusInternalServerError)
		return
	}

	utcNow := time.Now().UTC()
	localTime := utcNow.Add(time.Second * time.Duration(data.Timezone))

	response := FinalResponse{
		Name:      data.CityName,
		Temp:      data.Main.Kelvin,
		LocalTime: localTime.Format("2006-01-02 15:04:05"),
	}

	json.NewEncoder(w).Encode(response)
}

func getTemperature(city string) (*WeatherResponse, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load API key: %v", err)
	}

	apiURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, config.ApiKey)
	fmt.Println("API URL:", apiURL)

	res, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: status %d, body: %s", res.StatusCode, string(body))
	}

	var weatherData WeatherResponse
	if err := json.Unmarshal(body, &weatherData); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return &weatherData, nil
}

func main() {
	http.HandleFunc("/weather/", weatherHandler)

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
