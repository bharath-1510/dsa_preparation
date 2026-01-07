package strings_fundamentals

import "testing"

// Assume function signature: func LongestCommonPrefix(strs []string) string
func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		input []string
		want string
	}{
		{"example", []string{"flower","flow","flight"}, "fl"},
		{"no_common", []string{"dog","racecar","car"}, ""},
		{"empty", []string{}, ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LongestCommonPrefix(tc.input)
			if got != tc.want {
				t.Errorf("LongestCommonPrefix(%v) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
