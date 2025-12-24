package arraymedium

import "testing"

// Assume function signature: func SubarrayWithGivenProduct(arr []int, p int) (int, int)
func TestSubarrayWithGivenProduct(t *testing.T) {
	tests := []struct {
		name  string
		arr   []int
		p     int
		want  [2]int // -1,-1 if not found
	}{
		{"example", []int{2, 3, 4, 2, 6}, 12, [2]int{1, 2}},
		{"none", []int{1, 2, 3}, 100, [2]int{-1, -1}},
		{"single", []int{5}, 5, [2]int{0, 0}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			i, j := SubarrayWithGivenProduct(tc.arr, tc.p)
			got := [2]int{i, j}
			if got != tc.want {
				t.Errorf("SubarrayWithGivenProduct(%v, %d) = %v, want %v", tc.arr, tc.p, got, tc.want)
			}
		})
	}
}
