package strings_fundamentals

import "testing"

// Assume function signature: func ZFunction(s string) []int
func TestZFunction(t *testing.T) {
	tests := []struct {
		name string
		input string
		want []int
	}{
		{"example", "aabcaabx", []int{0,1,0,0,3,1,0,0}},
		{"single_char", "a", []int{0}},
		{"empty", "", []int{}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ZFunction(tc.input)
			if len(got) != len(tc.want) {
				t.Errorf("ZFunction(%q) = %v, want %v", tc.input, got, tc.want)
				return
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Errorf("ZFunction(%q) = %v, want %v", tc.input, got, tc.want)
					return
				}
			}
		})
	}
}
