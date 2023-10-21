package tokenprovider

import (
	"errors"
	cerr "go-clean-architecture/common/error"
	"net/http"
)

var (
	ErrEncodingToken = cerr.NewCustomError(
		http.StatusInternalServerError,
		errors.New("error encoding the token"),
		"error encoding the token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = cerr.NewCustomError(
		http.StatusForbidden,
		errors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken",
	)
)
