package strings_fundamentals

import "testing"

// Assume function signature: func ImplementStrStr(haystack, needle string) int
func TestImplementStrStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		{"example", "hello", "ll", 2},
		{"not_found", "aaaaa", "bba", -1},
		{"empty_needle", "abc", "", 0},
		{"empty_haystack", "", "a", -1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ImplementStrStr(tc.haystack, tc.needle)
			if got != tc.want {
				t.Errorf("ImplementStrStr(%q, %q) = %d, want %d", tc.haystack, tc.needle, got, tc.want)
			}
		})
	}
}
