package arraymedium

import "testing"

// Assume function signature: func SingleNumber(arr []int) int
func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"example", []int{2, 3, 2, 4, 3}, 4},
		{"all_unique", []int{7}, 7},
		{"longer", []int{1, 1, 2, 2, 3, 4, 4}, 3},
		{"negatives", []int{-1, -1, -2}, -2},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := SingleNumber(tc.arr)
			if got != tc.want {
				t.Errorf("SingleNumber(%v) = %d, want %d", tc.arr, got, tc.want)
			}
		})
	}
}
