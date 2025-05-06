package validators

import (
	"errors"
	"strings"

	"github.com/nyaruka/phonenumbers"
)

// ParseAndFormatPhone validates the phone number and returns its E.164 format.
func ParseAndFormatPhone(phone string, region string) (string, error) {
	phone = strings.TrimSpace(phone)
	if phone == "" {
		return "", errors.New("phone number is required")
	}

	if region == "" {
		region = "US" // fallback default
	}

	num, err := phonenumbers.Parse(phone, region)
	if err != nil {
		return "", errors.New("invalid phone number format")
	}

	if !phonenumbers.IsValidNumber(num) {
		return "", errors.New("invalid phone number")
	}

	formatted := phonenumbers.Format(num, phonenumbers.E164)
	return formatted, nil
}
