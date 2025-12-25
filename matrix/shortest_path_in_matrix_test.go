package matrix

import "testing"

// Assume function signature: func ShortestPathInMatrix(matrix [][]int) int
func TestShortestPathInMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{"example", [][]int{{0,1,0},{1,0,1},{0,0,0}}, 4},
		{"blocked", [][]int{{1,1},{1,1}}, -1},
		{"single", [][]int{{0}}, 1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ShortestPathInMatrix(tc.matrix)
			if got != tc.want {
				t.Errorf("ShortestPathInMatrix(%v) = %d, want %d", tc.matrix, got, tc.want)
			}
		})
	}
}
