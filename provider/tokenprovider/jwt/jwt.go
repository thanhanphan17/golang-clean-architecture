package jwt

import (
	"go-clean-architecture/provider/tokenprovider"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtProvider struct {
	secret string
}

var _ tokenprovider.Provider = (*jwtProvider)(nil)

func NewTokenJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload,
	expiry int) (*tokenprovider.Token, error) {
	// generate the JWT
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(
				time.Second * time.Duration(expiry)).Unix(), //Local() => UTC(): for standard
			IssuedAt: time.Now().Local().Unix(), //Local() => UTC(): for standard
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	// return the token
	return &tokenprovider.Token{
		Token:     myToken,
		Expiry:    expiry,
		CreatedAt: time.Now(),
	}, nil
}

func (j *jwtProvider) Validate(myToken string) (*tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(myToken, &myClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.secret), nil
		},
	)

	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	// validate the token
	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	// return the token
	return &claims.Payload, nil
}
