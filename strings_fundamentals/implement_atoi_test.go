package strings_fundamentals

import "testing"

// Assume function signature: func ImplementAtoi(s string) int
func TestImplementAtoi(t *testing.T) {
	tests := []struct {
		name string
		input string
		want int
	}{
		{"example", "   -42", -42},
		{"positive", "4193 with words", 4193},
		{"overflow", "91283472332", 2147483647},
		{"underflow", "-91283472332", -2147483648},
		{"empty", "", 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ImplementAtoi(tc.input)
			if got != tc.want {
				t.Errorf("ImplementAtoi(%q) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
