// Package random its a package to generate a string
package random

import (
	"math/rand"
	"time"
)

var chars string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Generate generates a new random string from a length
func Generate(length int) (finalString string) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < length; i++ {
		finalString += string(chars[r1.Intn(61)])
	}

	return finalString
}
