package matrix

import "testing"

// Assume function signature: func WordSearch(board [][]byte, word string) bool
func TestWordSearch(t *testing.T) {
	tests := []struct {
		name   string
		board  [][]byte
		word   string
		want   bool
	}{
		{"example", [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}, "ABCCED", true},
		{"not_found", [][]byte{{'A','B'},{'C','D'}}, "ABCD", false},
		{"single", [][]byte{{'A'}}, "A", true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := WordSearch(tc.board, tc.word)
			if got != tc.want {
				t.Errorf("WordSearch(%v, %q) = %v, want %v", tc.board, tc.word, got, tc.want)
			}
		})
	}
}
