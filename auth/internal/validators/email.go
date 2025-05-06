package validators

import (
	"errors"
	"net/mail"
	"strings"
)

// ValidateEmail checks if the provided email is in a valid format.
// Returns nil if valid, otherwise an error describing the issue.
func ValidateEmail(email string) (string, error) {
	email = strings.TrimSpace(email)

	if email == "" {
		return "", errors.New("email is required")
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return "", errors.New("invalid email format")
	}

	return email, nil
}
