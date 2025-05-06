package validators

import (
	"errors"
	"strings"
	"unicode"
)

// ValidatePassword enforces strong password rules.
// Returns nil if password is valid, otherwise a descriptive error.
func ValidatePassword(password string) (string, error) {

	const minLen = 8
	const maxLen = 64

	password = strings.TrimSpace(password)

	if len(password) < minLen || len(password) > maxLen {
		return "", errors.New("password must be between 8 and 64 characters")
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasDigit   bool
		hasSpecial bool
	)

	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return "", errors.New("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return "", errors.New("password must contain at least one lowercase letter")
	}
	if !hasDigit {
		return "", errors.New("password must contain at least one number")
	}
	if !hasSpecial {
		return "", errors.New("password must contain at least one special character")
	}

	return password, nil
}
