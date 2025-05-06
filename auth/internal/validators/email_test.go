package validators

// import "testing"

// func TestValidateEmail(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		input   string
// 		wantErr bool
// 	}{
// 		// Valid cases
// 		{"ValidEmail", "user@example.com", false},
// 		{"ValidEmailWithWhitespace", "  user@example.com  ", false},
// 		{"ValidSubdomain", "user@mail.example.com", false},

// 		// Invalid cases
// 		{"EmptyString", "", true},
// 		{"OnlyWhitespace", "    ", true},
// 		{"MissingAtSymbol", "userexample.com", true},
// 		{"MissingDomain", "user@", true},
// 		{"InvalidChars", "user@exa mple.com", true},
// 		{"DoubleAt", "user@@example.com", true},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := ValidateEmail(tt.input)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("ValidateEmail(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
// 			}
// 		})
// 	}
// }
