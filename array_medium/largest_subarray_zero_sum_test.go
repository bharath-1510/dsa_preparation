package arraymedium

import "testing"

// Assume function signature: func LargestSubarrayZeroSum(arr []int) int
func TestLargestSubarrayZeroSum(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"example", []int{15, -2, 2, -8, 1, 7}, 5},
		{"all_zero", []int{0, 0, 0, 0}, 4},
		{"none", []int{1, 2, 3}, 0},
		{"empty", []int{}, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LargestSubarrayZeroSum(tc.arr)
			if got != tc.want {
				t.Errorf("LargestSubarrayZeroSum(%v) = %d, want %d", tc.arr, got, tc.want)
			}
		})
	}
}
