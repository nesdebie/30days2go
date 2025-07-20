package main

import (
    "fmt"
    "os"
    "strings"
)

func writeFile(filename string, s string) error {
    data := []byte(s)
    err := os.WriteFile(filename, data, 0666)
    if err != nil {
        fmt.Printf("Some thing went wrong !! %v", err)
        os.Exit(1)
    }
    return err
}

func readTextFile(filename string) string {
    data, err := os.ReadFile(filename)
    if err != nil {
        fmt.Println("Error reading file:", err)
        os.Exit(1)
    }
    return string(data)
}

func countCharacters(text string) int {
    return len(text)
}

func countWords(text string) int {
    words := strings.Fields(text)
    return len(words)
}

func countLines(text string) int {
    lines := strings.Split(text, "\n")
    return len(lines)
}

func countFrequency(s []string) map[string]int {
    m := make(map[string]int)
    for _, word := range s {
        count, ok := m[word]
        if !ok {
            count = 1
        } else {
            count++
        }
        m[word] = count
    }
  return m
}

func printingFrequency(m map[string]int, s []string) {
    visited := make(map[string]bool)
    fmt.Println("Printing each word frequency:")
    for _, word := range s {
        if !visited[word] {
            fmt.Println(word + ": " + fmt.Sprint(m[word]) + " time(s)")
            visited[word] = true
        }
    }
    fmt.Println()
}

func main() {
    var filename string
    if len(os.Args) == 2 {
        filename = os.Args[1]
        _, err := os.Stat(filename)
        if err != nil {
            fmt.Println("Error: ", err)
            os.Exit(1)
        }
    } else {
        filename := "myText.txt"

        _, err := os.Stat(filename)
        if err == nil {
            os.Remove(filename)
        }
        writeFile(filename, strings.Join(os.Args[1:], " "))
    }
    
    text := readTextFile(filename)
    s := strings.Fields(text)

    characterCount := countCharacters(text)
    wordCount := countWords(text)
    lineCount := countLines(text)
    m := countFrequency(s)

    printingFrequency(m, s)

    fmt.Println()
    fmt.Println("Character count: " + fmt.Sprint(characterCount))
    fmt.Println("Word count: " + fmt.Sprint(wordCount))
    fmt.Println("Line count: " + fmt.Sprint(lineCount))
}