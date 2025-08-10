package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "sync"
    "time"
)

func downloadImage(url string, folder string) error {
    client := http.Client{
        Timeout: 10 * time.Second,
    }

    resp, err := client.Get(url)
    if err != nil {
        return fmt.Errorf("HTTP error: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return fmt.Errorf("status code %d", resp.StatusCode)
    }

    fileName := filepath.Base(url)
    filePath := filepath.Join(folder, fileName)

    outFile, err := os.Create(filePath)
    if err != nil {
        return fmt.Errorf("error creating file: %w", err)
    }
    defer outFile.Close()

    _, err = io.Copy(outFile, resp.Body)
    if err != nil {
        return fmt.Errorf("error writing file: %w", err)
    }

    return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + " <url1> <url2> ... <urlN>")
		fmt.Println("Example: " + os.Args[0] + " https://example.com/image1.jpg https://example.com/image2.png https://example.com/image3.gif")
		os.Exit(1)
	}

    urls := os.Args[1:]

    folder := "images"
    os.MkdirAll(folder, os.ModePerm)

    var waitingGroup sync.WaitGroup

    maxConcurrent := 5
    semaphore := make(chan struct{}, maxConcurrent)

    for _, url := range urls {
        waitingGroup.Add(1)
        go func(u string) {
            defer waitingGroup.Done()

            semaphore <- struct{}{}

            fmt.Println("Downloading:", u)
            err := downloadImage(u, folder)
            if err != nil {
                fmt.Printf("Error downloading %s : %v\n", u, err)
            } else {
                fmt.Printf("Sucessfully downloaded: %s\n", u)
            }

            <-semaphore
        }(url)
    }

    waitingGroup.Wait()
}
