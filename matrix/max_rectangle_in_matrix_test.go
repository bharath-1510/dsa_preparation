package matrix

import "testing"

// Assume function signature: func MaxRectangleInMatrix(matrix [][]int) int
func TestMaxRectangleInMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{"example", [][]int{{0, 1, 1, 0}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 0, 0}}, 8},
		{"all_zero", [][]int{{0, 0}, {0, 0}}, 0},
		{"all_one", [][]int{{1, 1}, {1, 1}}, 4},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MaxRectangleInMatrix(tc.matrix)
			if got != tc.want {
				t.Errorf("MaxRectangleInMatrix(%v) = %d, want %d", tc.matrix, got, tc.want)
			}
		})
	}
}
