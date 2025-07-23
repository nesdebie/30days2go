package main

import (
    "fmt"
    "os"
    "strconv"
    "time"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: " + os.Args[0] + " <first N prime numbers>")
        return
    }

    raw_n_prime_numbers := os.Args[1]
    n_prime_numbers, err := strconv.Atoi(raw_n_prime_numbers)
    if err != nil || n_prime_numbers <= 0 {
        fmt.Println("Error: Please provide a valid positive integer for the number of prime numbers.")
        return
    }
    
    fmt.Println("Finding first " + raw_n_prime_numbers + " prime numbers...")

    start := time.Now()
    fmt.Print("2 ")
    for i, count, num := 0, 1, 3; count < n_prime_numbers; num+=2 {
        isPrime := true
        for j := 3; j*j <= num; j+=2 {
            if num%j == 0 {
                isPrime = false
                break
            }
        }
        if isPrime {
            fmt.Print(num, " ")
            i++
            count++
        }
    }
    end := time.Now()
    fmt.Println()
    fmt.Println("Time taken to start: ", end.Sub(start))
}