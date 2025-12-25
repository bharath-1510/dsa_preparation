package matrix

import "testing"

// Assume function signature: func SpiralTraversal(matrix [][]int) []int
func TestSpiralTraversal(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  []int
	}{
		{"3x3", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{"1x4", [][]int{{1, 2, 3, 4}}, []int{1, 2, 3, 4}},
		{"4x1", [][]int{{1}, {2}, {3}, {4}}, []int{1, 2, 3, 4}},
		{"empty", [][]int{}, []int{}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := SpiralTraversal(tc.input)
			if len(got) != len(tc.want) {
				t.Errorf("SpiralTraversal(%v) = %v, want %v", tc.input, got, tc.want)
				return
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Errorf("SpiralTraversal(%v) = %v, want %v", tc.input, got, tc.want)
					return
				}
			}
		})
	}
}
