package strings_fundamentals

import "testing"

// Assume function signature: func ValidParentheses(s string) bool
func TestValidParentheses(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"example", "()[]{}", true},
		{"invalid", "(]", false},
		{"empty", "", true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ValidParentheses(tc.input)
			if got != tc.want {
				t.Errorf("ValidParentheses(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
