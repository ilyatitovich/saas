package validators

// import (
// 	"testing"
// )

// func TestValidatePassword(t *testing.T) {
// 	tests := []struct {
// 		pw       string
// 		expected bool
// 	}{
// 		{"Short1!", false},
// 		{"alllowercase1!", false},
// 		{"ALLUPPERCASE1!", false},
// 		{"NoDigits!", false},
// 		{"NoSpecial1", false},
// 		{"Valid1!", true},
// 	}

// 	for _, tt := range tests {
// 		err := ValidatePassword(tt.pw)
// 		if (err == nil) != tt.expected {
// 			t.Errorf("Password %q validity = %v, expected %v", tt.pw, err == nil, tt.expected)
// 		}
// 	}
// }
