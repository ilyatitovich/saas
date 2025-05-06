package validators

import "net/mail"

// Email validation
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
