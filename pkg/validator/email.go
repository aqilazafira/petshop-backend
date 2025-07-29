package validator

import (
	"regexp"
	"strings"
)

// EmailRegex defines the regular expression for valid email format
// This regex balances strictness with practical email format requirements
const EmailRegex = `^[a-zA-Z0-9]([a-zA-Z0-9._%+-]*[a-zA-Z0-9])?@[a-zA-Z0-9]([a-zA-Z0-9-]*[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9-]*[a-zA-Z0-9])?)*\.[a-zA-Z]{2,}$`

// IsValidEmail validates email format using regex
func IsValidEmail(email string) bool {
	// Trim whitespace
	email = strings.TrimSpace(email)

	// Check if email is empty
	if email == "" {
		return false
	}

	// Check email length (RFC 5321 recommends max 254 characters)
	if len(email) > 254 {
		return false
	}

	// Compile and test regex
	re := regexp.MustCompile(EmailRegex)
	return re.MatchString(email)
}

// ValidateEmailFormat validates email and returns error message if invalid
func ValidateEmailFormat(email string) (bool, string) {
	if !IsValidEmail(email) {
		return false, "Invalid email format. Please provide a valid email address"
	}
	return true, ""
}
