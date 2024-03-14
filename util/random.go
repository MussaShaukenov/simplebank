package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

// The RandomInt function generates a random integer within a specified range.
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// The RandomString function generates a random string of a specified length.
func RandomString(n int) string {
	var stringBuider strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		stringBuider.WriteByte(c)
	}

	return stringBuider.String()
}

// The RandomOwner function generates a random owner name.
func RandomOwner() string {
	return RandomString(6)
}

// The RandomMoney function generates a random amount of money.
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// The RandomCurrency function generates a random currency.
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "KZT"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomAmount() int64 {
	return RandomInt(0, 1000)
}
