package controller

import "testing"

func TestValidateEmail(t *testing.T) {
	validEmail := "test@example.com"
	invalidEmail := "invalidemail"

	if !ValidateEmail(validEmail) {
		t.Errorf("Expected %s to be valid email", validEmail)
	}

	if ValidateEmail(invalidEmail) {
		t.Errorf("Expected %s to be invalid email", invalidEmail)
	}
}
