package validator

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	testCases := []struct {
		email    string
		expected bool
		name     string
	}{
		// Valid emails
		{"user@example.com", true, "simple valid email"},
		{"test.email@domain.co.id", true, "email with dots"},
		{"user123@example-site.com", true, "email with numbers and hyphens"},
		{"first.last+tag@example.org", true, "email with plus sign"},
		{"user_name@example.net", true, "email with underscore"},
		{"admin@localhost.localdomain", true, "localhost domain"},

		// Invalid emails
		{"", false, "empty email"},
		{"   ", false, "whitespace only"},
		{"plainaddress", false, "no @ symbol"},
		{"@missingusername.com", false, "missing username"},
		{"username@.com", false, "missing domain"},
		{"username@com", false, "missing TLD"},
		{"username@domain.", false, "incomplete TLD"},
		{"username@", false, "missing domain completely"},
		{"username.domain.com", false, "no @ symbol"},
		{"username@domain..com", false, "double dots in domain"},
		{"username@@domain.com", false, "double @ symbol"},
		{"user name@domain.com", false, "space in username"},
		{"username@domain .com", false, "space in domain"},
		{"username@domain,com", false, "comma instead of dot"},
		{"username@domain.c", false, "TLD too short"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsValidEmail(tc.email)
			if result != tc.expected {
				t.Errorf("IsValidEmail(%q) = %v; expected %v", tc.email, result, tc.expected)
			}
		})
	}
}

func TestValidateEmailFormat(t *testing.T) {
	// Test valid email
	isValid, errMsg := ValidateEmailFormat("test@example.com")
	if !isValid {
		t.Errorf("ValidateEmailFormat for valid email should return true, got false")
	}
	if errMsg != "" {
		t.Errorf("ValidateEmailFormat for valid email should return empty error message, got: %s", errMsg)
	}

	// Test invalid email
	isValid, errMsg = ValidateEmailFormat("invalid-email")
	if isValid {
		t.Errorf("ValidateEmailFormat for invalid email should return false, got true")
	}
	if errMsg == "" {
		t.Errorf("ValidateEmailFormat for invalid email should return error message, got empty string")
	}

	expectedErrorMsg := "Invalid email format. Please provide a valid email address"
	if errMsg != expectedErrorMsg {
		t.Errorf("ValidateEmailFormat error message should be %q, got %q", expectedErrorMsg, errMsg)
	}
}

func TestEmailLengthValidation(t *testing.T) {
	// Test email that's too long (over 254 characters)
	longEmail := "a"
	for len(longEmail) < 250 {
		longEmail += "a"
	}
	longEmail += "@example.com" // This will make it over 254 characters

	isValid := IsValidEmail(longEmail)
	if isValid {
		t.Errorf("IsValidEmail should return false for email longer than 254 characters")
	}
}
