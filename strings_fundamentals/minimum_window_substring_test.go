package strings_fundamentals

import "testing"

// Assume function signature: func MinimumWindowSubstring(s, t string) string
func TestMinimumWindowSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{"example", "ADOBECODEBANC", "ABC", "BANC"},
		{"not_found", "a", "b", ""},
		{"empty_t", "abc", "", ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MinimumWindowSubstring(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("MinimumWindowSubstring(%q, %q) = %q, want %q", tc.s, tc.t, got, tc.want)
			}
		})
	}
}
