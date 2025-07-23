package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type ExchangeResponse struct {
	Success bool               `json:"success"`
	Query   map[string]any     `json:"query"`
	Info    map[string]any     `json:"info"`
	Result  float64            `json:"result"`
}

func getExchangeRate(from, to string, amount float64) (float64, error) {
	url := fmt.Sprintf("https://api.frankfurter.app/latest?amount=%f&from=%s&to=%s", amount, from, to)

	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("network error: %v", err)
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, fmt.Errorf("failed to parse response: %v", err)
	}

	ratesRaw, ok := data["rates"]
	if !ok {
		return 0, fmt.Errorf("invalid currency code: '%s' or '%s'", from, to)
	}

	rates, ok := ratesRaw.(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("unexpected response format")
	}

	val, ok := rates[to]
	if !ok {
		return 0, fmt.Errorf("currency '%s' not found in response", to)
	}

	result, ok := val.(float64)
	if !ok {
		return 0, fmt.Errorf("unexpected value type for conversion result")
	}

	return result, nil
}



func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <FROM_CURRENCY> <TO_CURRENCY> <AMOUNT>")
		return
	}

	from := strings.ToUpper(os.Args[1])
	to := strings.ToUpper(os.Args[2])
	amount, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil || amount <= 0 {
		fmt.Println("Invalid amount. Please provide a positive number.")
		return
	}

	converted, err := getExchangeRate(from, to, amount)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("%.2f %s = %.2f %s\n", amount, from, converted, to)
}
