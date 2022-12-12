package helpers

import (
	"math/rand"
	"strings"
)

const randomInt = "1234567890"

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomPrice() int {
	return RandomInt(5000, 15000)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(randomInt)

	for i := 0; i < n; i++ {
		c := randomInt[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
