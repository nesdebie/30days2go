package main

import (
    "fmt"
    "net/http"
    "sync"
    "time"
	"os"
)

type PingResult struct {
    URL      string
    Duration time.Duration
    Error    error
}


func ping(url string, waitGroup *sync.WaitGroup, results chan<- PingResult) {
    defer waitGroup.Done()
    start := time.Now()
    resp, err := http.Get(url)
    duration := time.Since(start)
    if resp != nil {
        resp.Body.Close()
    }
    results <- PingResult{URL: url, Duration: duration, Error: err}
}


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + "<url1> <url2> ... <url_N>")
		fmt.Println("Example: " + os.Args[0] + " https://google.com https://github.com/nesdebie" )
		os.Exit(1)
	}

    urls := os.Args[1:]

    var waitGroup sync.WaitGroup
    results := make(chan PingResult, len(urls))

    for _, url := range urls {
        waitGroup.Add(1)
        go ping(url, &waitGroup, results)
    }

    go func() {
        waitGroup.Wait()
        close(results)
    }()

    for result := range results {
        if result.Error != nil {
            fmt.Printf("Error for %s: %v\n", result.URL, result.Error)
        } else {
            fmt.Printf("Ping to %s : %v\n", result.URL, result.Duration)
        }
    }
}
