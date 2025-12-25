package matrix

import "testing"

// Assume function signature: func MatrixSearch(matrix [][]int, target int) bool
func TestMatrixSearch(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{"example", [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}, 3, true},
		{"not_found", [][]int{{1,2},{3,4}}, 5, false},
		{"single", [][]int{{1}}, 1, true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MatrixSearch(tc.matrix, tc.target)
			if got != tc.want {
				t.Errorf("MatrixSearch(%v, %d) = %v, want %v", tc.matrix, tc.target, got, tc.want)
			}
		})
	}
}
