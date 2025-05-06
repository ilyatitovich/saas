package validators

// func TestValidatePhone(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		input   string
// 		wantErr bool
// 	}{
// 		// Valid examples
// 		{"ValidE164", "+14155552671", false},
// 		{"ValidLocalUS", "4155552671", false},
// 		{"ValidWithSpaces", "415 555 2671", false},
// 		{"ValidWithDashes", "415-555-2671", false},
// 		{"ValidWithParentheses", "(415) 555-2671", false},

// 		// Invalid examples
// 		{"Empty", "", true},
// 		{"WhitespaceOnly", "   ", true},
// 		{"InvalidChars", "415-abc-xyz1", true},
// 		{"TooShort", "123", true},
// 		{"TooLong", "+999999999999999999", true},
// 		{"MissingNumber", "+", true},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := ValidatePhone(tt.input, )
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("ValidatePhone(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
// 			}
// 		})
// 	}
// }
