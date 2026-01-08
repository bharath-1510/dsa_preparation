package strings_fundamentals

import "testing"

// Assume function signature: func KMPAlgorithm(text, pattern string) []int
func TestKMPAlgorithm(t *testing.T) {
	tests := []struct {
		name string
		text string
		pattern string
		want []int
	}{
		{"example", "ababcabcabababd", "ababd", []int{10}},
		{"not_found", "abcdef", "gh", []int{}},
		{"empty_pattern", "abc", "", []int{}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := KMPAlgorithm(tc.text, tc.pattern)
			if len(got) != len(tc.want) {
				t.Errorf("KMPAlgorithm(%q, %q) = %v, want %v", tc.text, tc.pattern, got, tc.want)
				return
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Errorf("KMPAlgorithm(%q, %q) = %v, want %v", tc.text, tc.pattern, got, tc.want)
					return
				}
			}
		})
	}
}
