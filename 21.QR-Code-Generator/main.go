package main

import (
	"os"
	"fmt"
	"path/filepath"

    "github.com/skip2/go-qrcode"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: " + os.Args[0] + " <text_or_url>")
		os.Exit(1)
	}

    content := os.Args[1]

	codePath := "img/" + content

	extension := filepath.Ext(codePath)
	if extension != ".png" {
		codePath = codePath + ".png"
	}

    err := qrcode.WriteFile(content, qrcode.Medium, 256, codePath)
    if err != nil {
        panic(err)
    }
}
