package validators

import "github.com/nyaruka/phonenumbers"

// Phone validation (E.164 format assumed)
func IsValidPhone(phone string) bool {
	num, err := phonenumbers.Parse(phone, "")
	return err == nil && phonenumbers.IsValidNumber(num)
}
