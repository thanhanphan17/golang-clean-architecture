package utils

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

// PasswordValidator checks if the given password meets the required criteria.
func PasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	return len(password) >= 8 &&
		containsUppercase(password) &&
		containsLowercase(password) &&
		containsDigit(password) &&
		containsSpecialCharacter(password)
}

// containsUppercase checks if the password contains at least one uppercase letter.
func containsUppercase(password string) bool {
	for _, char := range password {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

// containsLowercase checks if the password contains at least one lowercase letter.
func containsLowercase(password string) bool {
	for _, char := range password {
		if unicode.IsLower(char) {
			return true
		}
	}
	return false
}

// containsDigit checks if the password contains at least one digit.
func containsDigit(password string) bool {
	for _, char := range password {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

// containsSpecialCharacter checks if the password contains at least one special character.
func containsSpecialCharacter(password string) bool {
	for _, char := range password {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return true
		}
	}
	return false
}
