package main

import (
	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

type PlayRequest struct {
	Player string `json:"player"`
}

type PlayResponse struct {
	Player    string `json:"player"`
	Computer  string `json:"computer"`
	Result    string `json:"result"`
}


func determineResult(player, computer string) string {
	if player == computer {
		return "Draw"
	}

	winMap := map[string]string{
		"rock":    "scissors",
		"paper":   "rock",
		"scissors":   "paper",
	}

	if winMap[player] == computer {
		return "Win"
	}
	return "Lose"
}


func handlePlay(w http.ResponseWriter, r *http.Request) {
	var req PlayRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	choices := []string{"rock", "paper", "scissors"}
	computer := choices[rand.Intn(len(choices))]

	result := determineResult(req.Player, computer)

	resp := PlayResponse{
		Player:   req.Player,
		Computer: computer,
		Result:   result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}


func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}


func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/play", handlePlay)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	println("Serveur sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}