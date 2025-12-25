package matrix

import "testing"

// Assume function signature: func FloodFill(matrix [][]int, sr, sc, newColor int) [][]int
func TestFloodFill(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		sr, sc int
		newColor int
		want   [][]int
	}{
		{"example", [][]int{{1,1,1},{1,1,0},{1,0,1}}, 1, 1, 2, [][]int{{2,2,2},{2,2,0},{2,0,1}}},
		{"all_zero", [][]int{{0,0},{0,0}}, 0, 0, 1, [][]int{{1,1},{1,1}}},
		{"all_one", [][]int{{1,1},{1,1}}, 0, 0, 2, [][]int{{2,2},{2,2}}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FloodFill(tc.matrix, tc.sr, tc.sc, tc.newColor)
			if len(got) != len(tc.want) {
				t.Errorf("FloodFill(%v, %d, %d, %d) = %v, want %v", tc.matrix, tc.sr, tc.sc, tc.newColor, got, tc.want)
				return
			}
			for i := range got {
				for j := range got[i] {
					if got[i][j] != tc.want[i][j] {
						t.Errorf("FloodFill(%v, %d, %d, %d) = %v, want %v", tc.matrix, tc.sr, tc.sc, tc.newColor, got, tc.want)
						return
					}
				}
			}
		})
	}
}
