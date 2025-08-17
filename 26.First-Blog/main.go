package main

import (
	"log"
	"database/sql"
	"net/http"

	"mini-blog/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS articles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		author TEXT,
		title TEXT,
		content TEXT,
		created DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ListArticles(db)(w, r)
		case http.MethodPost:
			handlers.PostArticle(db)(w, r)
		default:
			http.Error(w, "Méthode non supportée", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/articles/", handlers.DeleteArticle(db))


	log.Println("Serveur sur :8080")
	http.ListenAndServe(":8080", nil)
}