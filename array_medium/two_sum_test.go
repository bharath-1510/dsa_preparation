package arraymedium

import "testing"

// Assume function signature: func TwoSum(arr []int, target int) (int, int)
func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   [2]int // -1,-1 if not found
	}{
		{"example", []int{2, 7, 11, 15}, 9, [2]int{0, 1}},
		{"no_pair", []int{1, 2, 3}, 7, [2]int{-1, -1}},
		{"duplicates", []int{3, 3, 4}, 6, [2]int{0, 1}},
		{"negatives", []int{-1, -2, -3, -4, -5}, -8, [2]int{2, 4}},
		{"empty", []int{}, 5, [2]int{-1, -1}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			i, j := TwoSum(tc.arr, tc.target)
			got := [2]int{i, j}
			// Accept either order
			if (got != tc.want) && (got != [2]int{tc.want[1], tc.want[0]}) {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tc.arr, tc.target, got, tc.want)
			}
			// Optional: check indices are valid and distinct if not -1
			if i != -1 && j != -1 {
				if i == j || i < 0 || j < 0 || i >= len(tc.arr) || j >= len(tc.arr) {
					t.Errorf("TwoSum returned invalid indices: %d, %d", i, j)
				}
			}
		})
	}
}
