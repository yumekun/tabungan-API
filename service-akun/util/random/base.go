package random

import (
	"math/rand"
	"time"
)

const (
	alphabetChars   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numericChars    = "0123456789"
	allChars        = alphabetChars + numericChars
	emailExtensions = "com net org biz info edu"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func generateString(characters string, length int) string {
	var result string

	for i := 0; i < length; i++ {
		result += string(characters[r.Intn(len(characters))])
	}

	return result
}


func GenerateNumericString(length int) string {
	return generateString(numericChars, length)
}