package utils

import (
	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	validatorRequest := validator.New()

	validatorRequest.RegisterValidation("password", PasswordValidator)
	validatorRequest.RegisterValidation("vnphone", PhoneValidator)

	return validatorRequest
}
