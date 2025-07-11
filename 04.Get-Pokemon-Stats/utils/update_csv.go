package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type NameEntry struct {
	Name     string `json:"name"`
	Language struct {
		Name string `json:"name"`
	} `json:"language"`
}

type PokemonSpecies struct {
	Names []NameEntry `json:"names"`
}

func main() {
	const maxID = 1025

	output := [][]string{
		{"id", "en", "fr", "de", "es", "it"},
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for i := 1; i <= maxID; i++ {
		url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-species/%d", i)
		apiResponse, err := client.Get(url)
		if err != nil {
			fmt.Printf("Error on ID %d: %v\n", i, err)
			continue
		}

		if apiResponse.StatusCode != 200 {
			apiResponse.Body.Close()
			continue
		}

		var species PokemonSpecies
		err = json.NewDecoder(apiResponse.Body).Decode(&species)
		apiResponse.Body.Close()
		if err != nil {
			fmt.Printf("Error decoding JSON on ID %d: %v\n", i, err)
			continue
		}

		languageMap := map[string]string{"en": "", "fr": "", "de": "", "es": "", "it": ""}
		for _, entry := range species.Names {
			if _, ok := languageMap[entry.Language.Name]; ok {
				languageMap[entry.Language.Name] = entry.Name
			}
		}

		if languageMap["en"] != "" && languageMap["fr"] != "" && languageMap["de"] != "" && languageMap["es"] != "" && languageMap["it"] != "" {
			row := []string{
				fmt.Sprintf("%d", i),
				languageMap["en"],
				languageMap["fr"],
				languageMap["de"],
				languageMap["es"],
				languageMap["it"],
			}
			output = append(output, row)
			fmt.Println("[ADD] #", i, " - ", languageMap["en"])
		}

		time.Sleep(100 * time.Millisecond)
	}

	file, err := os.Create("pokemon_names_multilang.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(output)
	if err != nil {
		panic(err)
	}

	fmt.Println("pokemon_names_multilang.csv generated successfully!")
}
