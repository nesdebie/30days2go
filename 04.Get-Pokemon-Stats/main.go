package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"

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

func removeAccents(s string) string {
	t := norm.NFD.String(s)
	result := make([]rune, 0, len(t))
	for _, r := range t {
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		result = append(result, r)
	}
	return string(result)
}

func loadPokemonNameMap(path string) (map[string]int, error) {
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

	nameToID := make(map[string]int)
	for _, row := range records[1:] {
		id, err := strconv.Atoi(row[0])
		if err != nil {
			continue
		}
		for _, name := range row[1:] {
			key := removeAccents(strings.ToLower(name))
			nameToID[key] = id
		}
	}
	return nameToID, nil
}

func getPokemonIDFromAnyLanguageCSV(input string, nameToID map[string]int) (int, error) {
	key := removeAccents(strings.ToLower(input))
	if id, ok := nameToID[key]; ok {
		return id, nil
	}
	return 0, fmt.Errorf("Pokémon name '%s' not found", input)
}

func printPokemonStatsAndTypes(pkmn *Pokemon) {
	fmt.Printf("%s:\n", strings.Title(pkmn.Name))
	fmt.Print("Type: ")
	for i, t := range pkmn.Types {
		if i > 0 {
			fmt.Print("")
		}
		fmt.Print("[" + t.Type.Name + "]")
	}
	fmt.Println()
	for _, stat := range pkmn.Stats {
		fmt.Printf("  %s: %d\n", stat.StatInfo.Name, stat.BaseStat)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "<pokemon>")
		return
	}

	nameToID, err := loadPokemonNameMap("pokemon_names_multilang.csv")
	if err != nil {
		fmt.Println("Error loading CSV:", err)
		return
	}

	input := os.Args[1]

	id, err := getPokemonIDFromAnyLanguageCSV(input, nameToID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching Pokémon data:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var pkmn Pokemon
	json.Unmarshal(body, &pkmn)

	printPokemonStatsAndTypes(&pkmn)
}
