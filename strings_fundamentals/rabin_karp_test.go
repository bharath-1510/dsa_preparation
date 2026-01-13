package strings_fundamentals

import "testing"

// Assume function signature: func RabinKarp(text, pattern string) []int
func TestRabinKarp(t *testing.T) {
	tests := []struct {
		name string
		text string
		pattern string
		want []int
	}{
		{"example", "AABAACAADAABAABA", "AABA", []int{0,9,12}},
		{"not_found", "abcdef", "gh", []int{}},
		{"empty_pattern", "abc", "", []int{}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := RabinKarp(tc.text, tc.pattern)
			if len(got) != len(tc.want) {
				t.Errorf("RabinKarp(%q, %q) = %v, want %v", tc.text, tc.pattern, got, tc.want)
				return
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Errorf("RabinKarp(%q, %q) = %v, want %v", tc.text, tc.pattern, got, tc.want)
					return
				}
			}
		})
	}
}
