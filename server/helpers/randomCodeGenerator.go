package helpers

import (
	"math/rand"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

//this will never generate the same random number as we are creating new random source and taking the current time as Seed,so everytime it runs,the random number will be generated based on current time
//chatgpt said that even if i generate 1billion random numbers everysecond,i wont repeat any number for atleast 292 years demm
func GenerateRandomCode(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano())) 
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(code)
}