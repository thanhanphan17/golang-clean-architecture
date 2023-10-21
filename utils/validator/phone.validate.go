package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/nyaruka/phonenumbers"
)

// PhoneValidator validates if a phone number is valid for the given country code.
func VNPhoneValidator(fl validator.FieldLevel) bool {
	// Get the phone number from the field
	phoneNumber := fl.Field().String()

	// Parse the phone number using the country code "VN"
	parsedNumber, err := phonenumbers.Parse(phoneNumber, "VN")
	if err != nil {
		// Return false if there was an error parsing the phone number
		return false
	}

	// Check if the parsed number is valid
	isValid := phonenumbers.IsValidNumber(parsedNumber)

	// Return the result of the validation
	return isValid
}
