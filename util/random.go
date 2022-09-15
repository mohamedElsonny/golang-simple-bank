package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano()) // random number generator
}

// RandomInt generate a number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generate a string of length n
func RandomString(n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(len(alphabet))]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generate a random owner name
func RandomOwner() string {
	return RandomString(8)
}

// RandomBalance generate a balance
func RandomBalance() int64 {
	return RandomInt(1, 1000)
}

// RandomCurrency generate a currency between USD, EUR, CAD
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	return currencies[rand.Intn(len(currencies))]
}
