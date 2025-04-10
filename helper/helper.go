package helper

import (
	"math/rand"

	"github.com/rishavkumar7/url-shortner/constants"
)

func GenerateRandomString(n int) string {
	const letters = constants.ShortUrlCodeCharacterSet
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
