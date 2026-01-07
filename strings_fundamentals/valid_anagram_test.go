package strings_fundamentals

import "testing"

// Assume function signature: func ValidAnagram(s, t string) bool
func TestValidAnagram(t *testing.T) {
	tests := []struct {
		name string
		s string
		t string
		want bool
	}{
		{"example", "anagram", "nagaram", true},
		{"not_anagram", "rat", "car", false},
		{"empty", "", "", true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ValidAnagram(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("ValidAnagram(%q, %q) = %v, want %v", tc.s, tc.t, got, tc.want)
			}
		})
	}
}
