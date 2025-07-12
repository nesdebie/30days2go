package main

import (
	"crypto/rand"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"
	"math/big"

	"golang.org/x/text/unicode/norm"
)

type NameEntry struct {
	Name     string `json:"name"`
	Language struct {
		Name string `json:"name"`
	} `json:"language"`
}

type Stat struct {
	BaseStat int `json:"base_stat"`
	StatInfo struct {
		Name string `json:"name"`
	} `json:"stat"`
}

type TypeEntry struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}

type Pokemon struct {
	Name  string      `json:"name"`
	Stats []Stat      `json:"stats"`
	Types []TypeEntry `json:"types"`
}

type PokemonResponse struct {
	Count int `json:"count"`
}

func removeAccents(s string) string {
	decomposed := norm.NFD.String(s)
	var result []rune
	for _, r := range decomposed {
		if !unicode.Is(unicode.Mn, r) {
			result = append(result, r)
		}
	}
	return string(result)
}

func loadCsv(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	csvMap := make(map[string]int)
	for _, row := range records[1:] {
		id, err := strconv.Atoi(row[0])
		if err != nil {
			continue
		}
		for _, name := range row[1:] {
			key := removeAccents(strings.ToLower(name))
			csvMap[key] = id
		}
	}
	return csvMap, nil
}

func getPokemonIdFromCsv(name string, csvMap map[string]int) (int, error) {
	key := removeAccents(strings.ToLower(name))
	if id, ok := csvMap[key]; ok {
		return id, nil
	}
	return 0, fmt.Errorf("Pokémon name '%s' not found", name)
}

func getMaxId() int {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon")
	if err != nil {
		fmt.Println("Error fetching max ID:", err)
		return 0
	}
	defer resp.Body.Close()

	var apiResp PokemonResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		fmt.Println("Error decoding response:", err)
		return 0
	}

	low, high := 1, apiResp.Count
	maxValid := 0
	for low <= high {
		mid := (low + high) / 2
		url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", mid)
		resp, err := http.Get(url)
		if err != nil {
			break
		}
		resp.Body.Close()

		if resp.StatusCode == 200 {
			maxValid = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return maxValid
}

func secureRandInt(max int) (int, error) {
	if max <= 0 {
		return 0, errors.New("max must be > 0")
	}
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(nBig.Int64()) + 1, nil
}

func getRandomValidPokemonId(maxId int) (int, error) {
	for {
		randomId, err := secureRandInt(maxId)
		if err != nil {
			return 0, err
		}
		url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", randomId)
		resp, err := http.Get(url)
		if err != nil {
			return 0, err
		}
		resp.Body.Close()

		if resp.StatusCode == 200 {
			return randomId, nil
		}
	}
}

func fetchPokemonById(id int) (*Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Pokémon ID %d not found", id)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var p Pokemon
	if err := json.Unmarshal(body, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

func printPokemonStatsAndTypes(p *Pokemon) {
	fmt.Printf("%s:\n", strings.Title(p.Name))
	fmt.Print("Type: ")
	for i, t := range p.Types {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print("[" + t.Type.Name + "]")
	}
	fmt.Println()
	for _, stat := range p.Stats {
		fmt.Printf("  %s: %d\n", stat.StatInfo.Name, stat.BaseStat)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "<pokemon name | rand>")
		return
	}
	pokemonName := os.Args[1]

	csvMap, err := loadCsv("pokemon_names_multilang.csv")
	if err != nil {
		fmt.Println("Error loading CSV:", err)
		return
	}

	var id int
	if key := strings.ToLower(pokemonName); key == "rand" || key == "random" {
		maxId := getMaxId()
		if maxId == 0 {
			fmt.Println("Couldn't determine max ID.")
			return
		}
		id, err = getRandomValidPokemonId(maxId)
		if err != nil {
			fmt.Println("Error getting random Pokémon:", err)
			return
		}
	} else {
		id, err = getPokemonIdFromCsv(pokemonName, csvMap)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	pokemon, err := fetchPokemonById(id)
	if err != nil {
		fmt.Println("Error fetching Pokémon data:", err)
		return
	}

	printPokemonStatsAndTypes(pokemon)
}
