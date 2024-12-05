package main

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomCode(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano())) 
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(code)
}

func main() {
	randomCode := generateRandomCode(10) // Generate a 10-character random code
	fmt.Println("Generated Random Code:", randomCode)
}
