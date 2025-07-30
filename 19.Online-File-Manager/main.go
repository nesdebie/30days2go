package main

import (
	"os"
	"io"
	"strings"
    "net/http"
    "log"
)


func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.Write([]byte(`
        <h1>Gestionnaire de fichiers en ligne</h1>
        <ul>
            <li><a href="/upload">Uploader un fichier</a></li>
            <li><a href="/files">Voir les fichiers</a></li>
        </ul>
    `))
}



func downloadHandler(w http.ResponseWriter, r *http.Request) {
    filename := strings.TrimPrefix(r.URL.Path, "/download/")
    http.ServeFile(w, r, "./uploads/" + filename)
}



func listFilesHandler(w http.ResponseWriter, r *http.Request) {
    files, _ := os.ReadDir("./uploads")
    w.Write([]byte("<ul>"))
    for _, file := range files {
        w.Write([]byte("<li><a href=\"/download/" + file.Name() + "\">" + file.Name() + "</a></li>"))
    }
    w.Write([]byte("</ul>"))
}


func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.ServeFile(w, r, "static/upload.html")
        return
    }
    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer file.Close()
    
    dst, _ := os.Create("./uploads/" + handler.Filename)
    defer dst.Close()
    io.Copy(dst, file)
}


func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/upload", uploadHandler)
    http.HandleFunc("/files", listFilesHandler)
    http.HandleFunc("/download/", downloadHandler)

    log.Println("Serveur sur :8080")
    http.ListenAndServe(":8080", nil)
}
