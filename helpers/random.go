package helpers

import "math/rand"

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomPrice() int {
	return RandomInt(5000, 15000)
}
