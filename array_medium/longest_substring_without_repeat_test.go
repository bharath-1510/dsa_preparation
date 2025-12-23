
package arraymedium

import "testing"

// Assume function signature: func LongestSubstringWithoutRepeat(s string) int
func TestLongestSubstringWithoutRepeat(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"example", "abcabcbb", 3},
		{"all_unique", "abcdef", 6},
		{"all_repeat", "aaaaa", 1},
		{"empty", "", 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LongestSubstringWithoutRepeat(tc.s)
			if got != tc.want {
				t.Errorf("LongestSubstringWithoutRepeat(%q) = %d, want %d", tc.s, got, tc.want)
			}
		})
	}
}
