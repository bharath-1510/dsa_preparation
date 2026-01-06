package strings_fundamentals

import "testing"

// Assume function signature: func ValidPalindrome(s string) bool
func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"example", "A man, a plan, a canal: Panama", true},
		{"not_palindrome", "race a car", false},
		{"empty", "", true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ValidPalindrome(tc.input)
			if got != tc.want {
				t.Errorf("ValidPalindrome(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
