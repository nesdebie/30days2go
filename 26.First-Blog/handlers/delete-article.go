package handlers

import (
    "net/http"
    "database/sql"
    "strconv"
    "strings"
)

func DeleteArticle(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodDelete {
            http.Error(w, "Unauthorized method", http.StatusMethodNotAllowed)
            return
        }

        parts := strings.Split(r.URL.Path, "/")
        if len(parts) < 3 {
            http.Error(w, "missing ID", http.StatusBadRequest)
            return
        }
        idStr := parts[2]
        id, err := strconv.Atoi(idStr)
        if err != nil {
            http.Error(w, "invalid ID", http.StatusBadRequest)
            return
        }

        _, err = db.Exec("DELETE FROM articles WHERE id = ?", id)
        if err != nil {
            http.Error(w, "SQL Error", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Article deleted"))
    }
}
