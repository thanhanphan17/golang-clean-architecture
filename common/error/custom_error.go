package cerr

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func NewErrorResponse(code int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: code,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(code int, root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(code, root, msg, root.Error(), key)
	}

	return NewErrorResponse(code, errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrDB(err error) *AppError {
	return NewErrorResponse(
		http.StatusBadRequest,
		err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(
		http.StatusBadRequest,
		err, "invalid request", err.Error(), "ERROR_INVALID_REQUEST")
}

func ErrInternal(err error) *AppError {
	return NewErrorResponse(
		http.StatusBadRequest,
		err, "internal error", err.Error(), "ERROR_INVALID_REQUEST")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_DELETE_%s", entity),
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("Cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_CREATE_%s", entity),
	)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_GET_%s", entity),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_UPDATE_%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_DELETE_%s", entity),
	)
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("%s deleted", strings.ToLower(entity)),
		fmt.Sprintf("ERR_%s_DELETED", entity),
	)
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("%s already exists", strings.ToLower(entity)),
		fmt.Sprintf("ERR_%s_ALREADY_EXISTS", entity),
	)
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("ERR_%s_NOT_FOUND", entity),
	)
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		"you have no permission",
		"ERR_NO_PERMISSION",
	)
}

func ErrWrongAuthHeader(err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		"wrong authen header",
		"ERR_WRONG_AUTH_HEADER",
	)
}

func ErrRecordNotFound(err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		"record not found",
		"ERR_RECORD_NOT_FOUND",
	)
}
