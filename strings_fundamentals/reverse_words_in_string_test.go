package strings_fundamentals

import "testing"

// Assume function signature: func ReverseWordsInString(s string) string
func TestReverseWordsInString(t *testing.T) {
	tests := []struct {
		name string
		input string
		want string
	}{
		{"example", "the sky is blue", "blue is sky the"},
		{"single_word", "hello", "hello"},
		{"empty", "", ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ReverseWordsInString(tc.input)
			if got != tc.want {
				t.Errorf("ReverseWordsInString(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
