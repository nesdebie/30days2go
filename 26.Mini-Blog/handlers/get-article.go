package handlers

import (
	"net/http"
	"database/sql"
	"fmt"
)

func ListArticles(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, author, title, content, created FROM articles")
		if err != nil {
			http.Error(w, "Base Error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var author, title, content, created string
			err = rows.Scan(&id, &author, &title, &content, &created)
			if err != nil {
				http.Error(w, "Reading Error", http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "%d | %s | %s | %s | %s\n", id, author, title, content, created)
		}
	}
}