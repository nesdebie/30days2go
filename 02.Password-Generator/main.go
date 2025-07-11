package main

import (
	"fmt"
	"strconv"
	"crypto/rand" // REALLY random
	"math/big"
)

func generatePassword(length int) string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/~`"
    password := make([]byte, length)
    for i := range password {
        randomNumber, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
        if err != nil {
            return ""
        }
        password[i] = charset[randomNumber.Int64()]
    }
    return string(password)
}

func main() {
	fmt.Println("____Password_Generator____")
	fmt.Print("Please enter the required size: ")

	var input string
	fmt.Scanln(&input)
	passwordLength, err := strconv.Atoi(input)
	if err != nil || passwordLength <= 15 {
		fmt.Println("Invalid input. Set to defqult size of 16 characters.")
		passwordLength = 16
	}
	if passwordLength > 64 {
		fmt.Println("Warning: Length is too long, setting to 64 characters.")
		passwordLength = 64
	}
	fmt.Printf("Generating a password of %d characters\n", passwordLength)
	password := generatePassword(passwordLength)
	fmt.Println(password)
}
