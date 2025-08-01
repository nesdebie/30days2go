package main

import (
    "html/template"
    "net/http"
    "os"
    "path/filepath"
    "sort"
)

type FileInfo struct {
    Name  string
    IsDir bool
}

func fileExplorerHandler(w http.ResponseWriter, r *http.Request) {
    dir := r.URL.Query().Get("dir")
    if dir == "" {
        dir = "."
    }
    files, err := os.ReadDir(dir)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    var fileInfos []FileInfo
    for _, f := range files {
        fileInfos = append(fileInfos, FileInfo{f.Name(), f.IsDir()})
    }
    sort.Slice(fileInfos, func(i, j int) bool { return fileInfos[i].Name < fileInfos[j].Name })

    tmpl := `<html>
        <head><title>Explorateur de fichiers</title></head>
        <body>
        <h1>Contenu du dossier: {{.Dir}}</h1>
        <ul>
            {{if ne .Dir "."}}
                <li><a href="/?dir={{.Parent}}">.. (parent)</a></li>
            {{end}}
            {{range .Files}}
                <li>
                    {{if .IsDir}}
                        <a href="/?dir={{$.Dir}}/{{.Name}}">{{.Name}}/</a>
                    {{else}}
                        {{.Name}}
                    {{end}}
                </li>
            {{end}}
        </ul>
        </body>
    </html>`
    parent := filepath.Dir(dir)
    t := template.Must(template.New("explorer").Parse(tmpl))
    t.Execute(w, map[string]interface{}{
        "Dir":    dir,
        "Files":  fileInfos,
        "Parent": parent,
    })
}

func main() {
    http.HandleFunc("/", fileExplorerHandler)
    http.ListenAndServe(":8080", nil)
}
