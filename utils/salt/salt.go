package utils

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// randSequence generates a random sequence of characters of length n.
func randSequence(n int) string {
	// Create a slice to store the random sequence
	b := make([]rune, n)

	// Create a new random number generator with a seed based on the current time
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random character for each index in the slice
	for i := range b {
		// Select a random index from the "letters" slice and assign it to the current index in the sequence slice
		b[i] = letters[r1.Intn(len(letters))]
	}

	// Convert the slice to a string and return it
	return string(b)
}

// GenSalt generates a random salt string of specified length.
// If length is less than 0, default length of 50 is used.
func GenSalt(length ...int) string {
	// Check if length is less than 0
	if len(length) == 0 {
		// Set default length to 50
		length = []int{50}
	}

	// Generate a random sequence of characters of specified length
	return randSequence(length[0])
}
