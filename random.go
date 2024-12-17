package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

// generate random decimal value to represent user Balance in test case
func RandomFloatAmount(min, max float64) string {
	randFloat := (min + max) * rand.Float64()
	return fmt.Sprintf("%.2f", randFloat)
}

const Alphabet = "abcdefghijklmnopqrstuvwxyz"

// generate random string of char
func RandomString(nChar int) string {
	var sb strings.Builder

	lenAlphabet := len(Alphabet)

	for i := 0; i < nChar; i++ {
		char := Alphabet[rand.Intn(lenAlphabet)]

		sb.WriteByte(char)
	}

	return sb.String()
}

// generate random name for test cases
func GenerateOwnerName() string {
	return RandomString(6)
}

// generate random currency user currency for test
func GenerateRandomCurrency() string {
	currencyArray := [4]string{"USD", "NAIRA", "CAD", "EUR"}
	n := len(currencyArray)
	return currencyArray[rand.Intn(n)]
}

// generate random amounts for test cases
func GenerateRandomAmount() string {
	return RandomFloatAmount(0, 100000)
}
