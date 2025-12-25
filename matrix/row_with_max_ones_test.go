package matrix

import "testing"

// Assume function signature: func RowWithMaxOnes(matrix [][]int) int
func TestRowWithMaxOnes(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{"example", [][]int{{0,1,1,1},{0,0,1,1},{1,1,1,1},{0,0,0,0}}, 2},
		{"all_zero", [][]int{{0,0},{0,0}}, 0},
		{"all_one", [][]int{{1,1},{1,1}}, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := RowWithMaxOnes(tc.matrix)
			if got != tc.want {
				t.Errorf("RowWithMaxOnes(%v) = %d, want %d", tc.matrix, got, tc.want)
			}
		})
	}
}
