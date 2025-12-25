package matrix

import "testing"

// Assume function signature: func RotateImage(matrix [][]int) [][]int
func TestRotateImage(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		want   [][]int
	}{
		{"3x3", [][]int{{1,2,3},{4,5,6},{7,8,9}}, [][]int{{7,4,1},{8,5,2},{9,6,3}}},
		{"1x1", [][]int{{5}}, [][]int{{5}}},
		{"2x2", [][]int{{1,2},{3,4}}, [][]int{{3,1},{4,2}}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := RotateImage(tc.input)
			if len(got) != len(tc.want) {
				t.Errorf("RotateImage(%v) = %v, want %v", tc.input, got, tc.want)
				return
			}
			for i := range got {
				for j := range got[i] {
					if got[i][j] != tc.want[i][j] {
						t.Errorf("RotateImage(%v) = %v, want %v", tc.input, got, tc.want)
						return
					}
				}
			}
		})
	}
}
