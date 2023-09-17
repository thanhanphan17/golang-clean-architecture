package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/nyaruka/phonenumbers"
)

func PhoneValidator(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()

	parsedNumber, err := phonenumbers.Parse(phoneNumber, "VN")
	if err != nil {
		return false
	}

	isValid := phonenumbers.IsValidNumber(parsedNumber)

	return isValid
}
