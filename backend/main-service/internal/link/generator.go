// Package link provides URL shortening functionality
package link

import (
	"math/rand"
)

func generatePseudoLink() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	key := make([]byte, 6)
	for i := range key {
		key[i] = charset[rand.Intn(len(charset))]
	}
	return "http://localhost:8080/" + string(key)
}

// "https://short.ly/" +
