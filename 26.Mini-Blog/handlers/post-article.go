package handlers

import (
    "net/http"
    "database/sql"
    "encoding/json"
    "mini-blog/models"
)

func PostArticle(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Unauthorized method", http.StatusMethodNotAllowed)
            return
        }
        var input models.Article
        err := json.NewDecoder(r.Body).Decode(&input)
        if err != nil {
            http.Error(w, "invalid JSON ", http.StatusBadRequest)
            return
        }
        _, err = db.Exec(
            "INSERT INTO articles (author, title, content) VALUES (?, ?, ?)",
            input.Author, input.Title, input.Content,
        )
        if err != nil {
            http.Error(w, "SQL Error", http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
        w.Write([]byte("Article created"))
    }
}
