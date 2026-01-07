package strings_fundamentals

import "testing"

// Assume function signature: func LongestPalindromicSubstring(s string) string
func TestLongestPalindromicSubstring(t *testing.T) {
	tests := []struct {
		name string
		input string
		want1 string
		want2 string
	}{
		{"example", "babad", "bab", "aba"},
		{"single_char", "a", "a", "a"},
		{"empty", "", "", ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LongestPalindromicSubstring(tc.input)
			if got != tc.want1 && got != tc.want2 {
				t.Errorf("LongestPalindromicSubstring(%q) = %q, want %q or %q", tc.input, got, tc.want1, tc.want2)
			}
		})
	}
}
