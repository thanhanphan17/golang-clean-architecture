package biz

import (
	cerr "go-clean-architecture/common/error"
	"net/http"
)

func ErrPasswordIsNotValid(err error) *cerr.AppError {
	return cerr.NewCustomError(
		http.StatusBadRequest,
		err,
		"password must have from 8 to 30 characters",
		"ErrWrongPassword",
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

func ErrEmailHasExisted(err error) *cerr.AppError {
	return cerr.NewCustomError(
		http.StatusBadRequest,
		err,
		"email has existed",
		"ErrExistedEmail",
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

func ErrPhoneHasNotVerified(err error) *cerr.AppError {
	return cerr.NewCustomError(
		http.StatusBadRequest,
		err,
		"phone has not been verified",
		"ErrPhoneHasNotVerified",
	)
}

func ErrPhoneHasExisted(err error) *cerr.AppError {
	return cerr.NewCustomError(
		http.StatusBadRequest,
		err,
		"ErrPhoneHasExisted",
		"ErrPhoneHasExisted",
	)
}
