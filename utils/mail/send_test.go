package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	_ = os.Setenv("CONFIG_PATH", "../../config/env")
	_ = os.Setenv("CONFIG_ENV", "local")

	// Test case: all parameters are valid
	t.Run("all parameters are valid", func(t *testing.T) {
		err := Send("thanhanphan17@gmail.com", "OTP", "./template/otp.html", map[string]interface{}{
			"CustomerName": "Phan Thanh An",
			"OTP":          "123456",
		})

		assert.NoError(t, err)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
	})

	// Test case: to parameter is empty
	t.Run("to parameter is empty", func(t *testing.T) {
		err := Send("", "Test Subject", "./template/otp.html", map[string]interface{}{})
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	})

	// Test case: subject parameter is empty
	t.Run("subject parameter is empty", func(t *testing.T) {
		err := Send("test@example.com", "", "./template/otp.html", map[string]interface{}{})
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	})

	// Test case: templatePath parameter is empty
	t.Run("templatePath parameter is empty", func(t *testing.T) {
		err := Send("test@example.com", "Test Subject", "", map[string]interface{}{})
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	})

	// Test case: all parameters are empty
	t.Run("all parameters are empty", func(t *testing.T) {
		err := Send("", "", "", map[string]interface{}{})
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	})
}
