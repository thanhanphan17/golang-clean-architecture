package utils

import "testing"

func TestGenSalt(t *testing.T) {
	// Test case: length is positive
	t.Run("Positive Length", func(t *testing.T) {
		length := 10
		salt := GenSalt(length)
		if len(salt) != length {
			t.Errorf("Expected salt length: %d, but got: %d", length, len(salt))
		}
	})

	// Test case: length is 0
	t.Run("Zero Length", func(t *testing.T) {
		length := 0
		salt := GenSalt(length)
		if len(salt) != 50 {
			t.Errorf("Expected salt length: 50, but got: %d", len(salt))
		}
	})

	// Test case: length is negative
	t.Run("Negative Length", func(t *testing.T) {
		length := -10
		salt := GenSalt(length)
		if len(salt) != 50 {
			t.Errorf("Expected salt length: 50, but got: %d", len(salt))
		}
	})
}
