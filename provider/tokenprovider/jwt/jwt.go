package jwt

import (
	"encoding/json"
	"go-clean-architecture/provider/tokenprovider"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtProvider struct {
	secret string
}

// NewJWTProvider creates a new instance of the TokenJWTProvider struct.
// It takes a secret string as a parameter and returns a pointer to the Provider struct.
func NewJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

// Generate generates a JWT token with the given data and expiry.
// It returns the generated token and any error encountered.
func (j *jwtProvider) Generate(
	payload map[string]interface{}, // The data for creating the token
	expiry uint, // The expiry time in seconds
) (map[string]interface{}, error) {
	payloadStruct, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	data := tokenprovider.TokenPayload{}
	if err := json.Unmarshal(payloadStruct, &data); err != nil {
		return nil, err
	}

	// Generate the JWT token
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": data.UserID,                                                // Set the user id in the token
		"role":    data.Role,                                                  // Set the role in the token
		"type":    data.Type,                                                  // Set the type in the token
		"exp":     time.Now().Add(time.Second * time.Duration(expiry)).Unix(), // Set the expiry time
		"iat":     time.Now().Unix(),                                          // Set the token creation time
	})

	// Sign the token with the secret key
	tokenString, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, tokenprovider.ErrEncodingToken
	}

	// Create a token object with the generated token, expiry, and creation time
	token := map[string]interface{}{
		"token":      tokenString,
		"expiry":     expiry,
		"created_at": time.Now(),
	}

	return token, nil
}

// Validate validates a token and returns the token payload.
//
// Parameters:
// - tokenString: The token string to be validated.
//
// Returns:
// - TokenPayload: The token payload if the token is valid.
// - error: An error if the token is invalid.
func (j *jwtProvider) Validate(tokenString string) (*tokenprovider.TokenPayload, error) {
	// Parse the token with the secret key.
	secretKey := []byte(j.secret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		slog.Info("JWT" + err.Error())
		return nil, tokenprovider.ErrInvalidToken
	}

	// Check if the token is valid.
	if token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, tokenprovider.ErrInvalidToken
		}

		return &tokenprovider.TokenPayload{
			UserID: claims["user_id"].(string),
			Role:   claims["role"].(string),
			Type:   claims["type"].(string),
		}, nil
	}

	return nil, tokenprovider.ErrInvalidToken
}
