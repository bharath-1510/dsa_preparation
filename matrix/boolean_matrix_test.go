package matrix

import "testing"

// Assume function signature: func BooleanMatrix(matrix [][]int) [][]int
func TestBooleanMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{"example", [][]int{{1,0},{0,0}}, [][]int{{1,1},{1,0}}},
		{"all_zero", [][]int{{0,0},{0,0}}, [][]int{{0,0},{0,0}}},
		{"all_one", [][]int{{1,1},{1,1}}, [][]int{{1,1},{1,1}}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := BooleanMatrix(tc.matrix)
			if len(got) != len(tc.want) {
				t.Errorf("BooleanMatrix(%v) = %v, want %v", tc.matrix, got, tc.want)
				return
			}
			for i := range got {
				for j := range got[i] {
					if got[i][j] != tc.want[i][j] {
						t.Errorf("BooleanMatrix(%v) = %v, want %v", tc.matrix, got, tc.want)
						return
					}
				}
			}
		})
	}
}
