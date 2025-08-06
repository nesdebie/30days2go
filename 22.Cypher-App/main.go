package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/term"
	"syscall"
)

func encryptAESGCM(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

func decryptAESGCM(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce := ciphertext[:nonceSize]
	cipherData := ciphertext[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}


func askKey() ([]byte, error) {
	fmt.Print("Enter encryption key (16, 24, or 32 characters): ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return nil, err
	}

	key := []byte(strings.TrimSpace(string(bytePassword)))
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, fmt.Errorf("key must be 16, 24, or 32 bytes (got %d)", len(key))
	}

	return key, nil
}


func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	input := os.Args[1]
	stat, err := os.Stat(input)

	if err != nil || stat.IsDir() {
		fmt.Println("Invalid file.")
		os.Exit(1)
	}

	key, err := askKey()
	if err != nil {
		fmt.Println("Key error:", err)
		os.Exit(1)
	}

	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(2)
	}

	if strings.HasSuffix(input, ".encrypted") {
		fmt.Println("Decrypting:", input)
		plaintext, err := decryptAESGCM(data, key)
		if err != nil {
			fmt.Println("Decryption error:", err)
			os.Exit(3)
		}

		outputFile := strings.TrimSuffix(input, ".encrypted")
		err = os.WriteFile(outputFile, plaintext, 0644)
		if err != nil {
			fmt.Println("Error writing decrypted file:", err)
			os.Exit(4)
		}

		err = os.Remove(input)
		if err == nil {
			fmt.Println("Decrypted and replaced:", outputFile)
		} else {
			fmt.Println("Decrypted, but could not delete original:", err)
		}

	} else {
		fmt.Println("Encrypting:", input)
		ciphertext, err := encryptAESGCM(data, key)
		if err != nil {
			fmt.Println("Encryption error:", err)
			os.Exit(3)
		}

		outputFile := input + ".encrypted"
		err = os.WriteFile(outputFile, ciphertext, 0644)
		if err != nil {
			fmt.Println("Error writing encrypted file:", err)
			os.Exit(4)
		}

		err = os.Remove(input)
		if err != nil {
			fmt.Println("Encrypted, but could not delete original:", err)
		}
		fmt.Println("Encrypted and replaced with:", outputFile)
	}
}
