package main

import (
    "fmt"
    "time"

    "simple-cache/cache"
)

func main() {
    c := cache.NewCache()
    c.Set("foo", "bar", 2*time.Second)

    value, found := c.Get("foo")
    fmt.Println("Found before expiration:", found, value)

    time.Sleep(3 * time.Second)

    value, found = c.Get("foo")
    fmt.Println("Found after expiration:", found, value)
}
