package business

import (
	cerr "go-clean-architecture/common/error"
	"net/http"
)

func ErrInvalidInfo(err error) *cerr.AppError {
	return cerr.NewCustomError(
		http.StatusBadRequest,
		err,
		"email or password not correct",
		"ErrInvalidInfo",
	)
}

func ErrEmailIsNotValid(err error) *cerr.AppError {
	return cerr.NewCustomError(
		http.StatusBadRequest,
		err,
		"email is not valid",
		"ErrInvalidEmail",
	)
}

func ErrEmailIsNotExisted(err error) *cerr.AppError {
	return cerr.NewCustomError(
		http.StatusBadRequest,
		err,
		"email is not existed",
		"ErrNotExistedEmail",
	)
}

func ErrEmailHasExisted(err error) *cerr.AppError {
	return cerr.NewCustomError(
		http.StatusBadRequest,
		err,
		"email has existed",
		"ErrExistedEmail",
	)
}

func ErrEmailNotVerified(err error) *cerr.AppError {
	return cerr.NewCustomError(
		http.StatusBadRequest,
		err,
		"email has not been verified",
		"ErrEmailHasNotVerified",
	)
}

func ErrEmailAlreadyVerified(err error) *cerr.AppError {
	return cerr.NewCustomError(
		http.StatusBadRequest,
		err,
		"email has been verified already",
		"ErrEmailAlreadyVerified",
	)
}

func ErrFullNameIsEmpty(err error) *cerr.AppError {
	return cerr.NewCustomError(
		http.StatusBadRequest,
		err,
		"full name can not be blank",
		"ErrBlankFullName",
	)
}

func ErrFullNameTooLong(err error) *cerr.AppError {
	return cerr.NewCustomError(
		http.StatusBadRequest,
		err,
		"full name too long, max character is 30",
		"ErrTooLongFullName",
	)
}

func ErrInvalidOTP(err error) *cerr.AppError {
	return cerr.NewCustomError(
		http.StatusBadRequest,
		err,
		"otp is not valid",
		"ErrInvalidOTP",
	)
}
