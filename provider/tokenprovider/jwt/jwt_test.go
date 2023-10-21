package jwt

import (
	"go-clean-architecture/provider/tokenprovider"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestValidateToken(t *testing.T) {
	provider := &jwtProvider{secret: "secret"}

	t.Run("Valid token", func(t *testing.T) {
		// Create a valid token
		claims := jwt.MapClaims{
			"user_id": "123",
			"role":    "admin",
			"type":    "access",
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString([]byte("secret"))

		// Call the Validate function
		payload, err := provider.Validate(tokenString)

		// Check if the payload and error are as expected
		assert.NoError(t, err)
		assert.Equal(t, "123", payload.UserID)
		assert.Equal(t, "admin", payload.Role)
		assert.Equal(t, "access", payload.Type)
	})

	t.Run("Invalid token", func(t *testing.T) {
		// Create an invalid token
		tokenString := "invalid_token"

		// Call the Validate function
		payload, err := provider.Validate(tokenString)

		// Check if the error is as expected
		assert.Equal(t, tokenprovider.ErrInvalidToken, err)
		assert.Nil(t, payload)
	})
}

func TestGenerateToken(t *testing.T) {
	provider := &jwtProvider{secret: "mysecret"}

	// Test case 1: Generate token with valid data and expiry
	t.Run("Generate token with valid data and expiry", func(t *testing.T) {
		payload := map[string]interface{}{
			"user_id": "user1",
			"role":    "admin",
			"type":    "access",
		}
		expiry := uint(3600)

		token, err := provider.Generate(payload, expiry)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		// Verify the generated token
		if token["token"] == "" {
			t.Errorf("Expected non-empty token, got empty token")
		}
		if token["expiry"].(uint) != expiry {
			t.Errorf("Expected expiry %d, got %d", expiry, token["expiry"])
		}
		if token["created_at"].(time.Time).IsZero() {
			t.Errorf("Expected non-zero creation time, got zero time")
		}
	})

	// Test case 2: Generate token with empty data and expiry
	t.Run("Generate token with empty data and expiry", func(t *testing.T) {
		payload := map[string]interface{}{}
		expiry := uint(0)

		token, err := provider.Generate(payload, expiry)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		// Verify the generated token
		if token["token"] == "" {
			t.Errorf("Expected non-empty token, got empty token")
		}
		if token["expiry"].(uint) != expiry {
			t.Errorf("Expected expiry %d, got %d", expiry, token["expiry"])
		}
		if token["created_at"].(time.Time).IsZero() {
			t.Errorf("Expected non-zero creation time, got zero time")
		}
	})
}
