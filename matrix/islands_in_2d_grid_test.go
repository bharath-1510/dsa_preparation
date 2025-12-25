package matrix

import "testing"

// Assume function signature: func IslandsIn2DGrid(matrix [][]int) int
func TestIslandsIn2DGrid(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{"example", [][]int{{1,1,0,0,0},{1,1,0,0,0},{0,0,1,0,0},{0,0,0,1,1}}, 3},
		{"all_zero", [][]int{{0,0},{0,0}}, 0},
		{"all_one", [][]int{{1,1},{1,1}}, 1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IslandsIn2DGrid(tc.matrix)
			if got != tc.want {
				t.Errorf("IslandsIn2DGrid(%v) = %d, want %d", tc.matrix, got, tc.want)
			}
		})
	}
}
