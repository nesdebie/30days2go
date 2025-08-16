package main

import (
	"fmt"
    "os"
    "bufio"
    "regexp"
    "strings"
    "sort"
    "log"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: " + os.Args[0] + " <logfile>")
		os.Exit(1)
	}
	logfile, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer logfile.Close()

    scanner := bufio.NewScanner(logfile)
    stats := make(map[string]int)

    errorRegexp := regexp.MustCompile(`ERROR`)
    var errorLines []string

    for scanner.Scan() {
        line := scanner.Text()

        date := strings.Split(line, " ")[0]
        stats[date]++

        if errorRegexp.MatchString(line) {
            errorLines = append(errorLines, line)
        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    type stat struct { Date string; Count int }
    var countSlice []stat
    for d, c := range stats {
        countSlice = append(countSlice, stat{Date: d, Count: c})
    }
    sort.Slice(countSlice, func(i, j int) bool { return countSlice[i].Count > countSlice[j].Count })

    log.Println("Stats per date:")
    for _, slice:= range countSlice {
        log.Printf("%s: %d\n", slice.Date, slice.Count)
    }
    log.Printf("Number of ERROR lines: %d\n", len(errorLines))
}
