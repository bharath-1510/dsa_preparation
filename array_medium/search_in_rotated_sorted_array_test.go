package arraymedium

import "testing"

// Assume function signature: func SearchInRotatedSortedArray(arr []int, target int) int
func TestSearchInRotatedSortedArray(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   int
	}{
		{"example", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{"not_found", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{"single", []int{1}, 1, 0},
		{"empty", []int{}, 1, -1},
		{"no_rotation", []int{1, 2, 3, 4, 5}, 3, 2},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := SearchInRotatedSortedArray(tc.arr, tc.target)
			if got != tc.want {
				t.Errorf("SearchInRotatedSortedArray(%v, %d) = %d, want %d", tc.arr, tc.target, got, tc.want)
			}
		})
	}
}
