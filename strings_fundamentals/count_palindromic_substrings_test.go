package strings_fundamentals

import "testing"

// Assume function signature: func CountPalindromicSubstrings(s string) int
func TestCountPalindromicSubstrings(t *testing.T) {
	tests := []struct {
		name string
		input string
		want int
	}{
		{"example", "aaa", 6},
		{"single_char", "a", 1},
		{"empty", "", 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CountPalindromicSubstrings(tc.input)
			if got != tc.want {
				t.Errorf("CountPalindromicSubstrings(%q) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
