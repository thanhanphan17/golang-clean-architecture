package tokenprovider

import (
	"errors"
	cerr "go-clean-architecture/common/error"
	"net/http"
)

var (
	ErrNotFound = cerr.NewCustomError(
		http.StatusBadRequest,
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
	)

	ErrEncodingToken = cerr.NewCustomError(
		http.StatusBadRequest,
		errors.New("error encoding the token"),
		"error encoding the token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = cerr.NewCustomError(
		http.StatusBadRequest,
		errors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken",
	)
)
