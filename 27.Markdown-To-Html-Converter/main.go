package main

import (
    "fmt"
	"os"
    "regexp"
    "strings"
	"io/ioutil"
)

func MarkdownToHTML(markdown string) string {
    reH1 := regexp.MustCompile(`(?m)^# (.+)$`)
    html := reH1.ReplaceAllString(markdown, "<h1>$1</h1>")
    
    reH2 := regexp.MustCompile(`(?m)^## (.+)$`)
    html = reH2.ReplaceAllString(html, "<h2>$1</h2>")

    reBold := regexp.MustCompile(`\*\*(.+?)\*\*`)
    html = reBold.ReplaceAllString(html, "<strong>$1</strong>")

    reItalics := regexp.MustCompile(`\*(.+?)\*`)
    html = reItalics.ReplaceAllString(html, "<em>$1</em>")

    lignes := strings.Split(html, "\n")
    for i, ligne := range lignes {
        if !strings.HasPrefix(ligne, "<h") && strings.TrimSpace(ligne) != "" {
            lignes[i] = "<p>" + ligne + "</p>"
        }
    }
    return strings.Join(lignes, "\n")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: " + os.Args[0] + " example.md")
		os.Exit(1)
	}
	filepath, err := os.Open(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	defer filepath.Close()
	mdfile, err := filepath.Stat()
	if err != nil || mdfile.IsDir() {
		fmt.Println("Error: file invalid or not found")
		os.Exit(1)
	}

	content, err := ioutil.ReadAll(filepath)
    if err != nil {
        fmt.Println("Error: cannot read file")
        os.Exit(1)
    }
    result := MarkdownToHTML(string(content))
    fmt.Println(result)
}
