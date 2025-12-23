package arraymedium

import "testing"

// Assume function signature: func LongestSubarraySumK(arr []int, k int) int
func TestLongestSubarraySumK(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		k    int
		want int
	}{
		{"example", []int{10, 5, 2, 7, 1}, 15, 4},
		{"none", []int{1, 2, 3}, 100, 0},
		{"negatives", []int{1, -1, 5, -2, 3}, 3, 4},
		{"empty", []int{}, 10, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LongestSubarraySumK(tc.arr, tc.k)
			if got != tc.want {
				t.Errorf("LongestSubarraySumK(%v, %d) = %d, want %d", tc.arr, tc.k, got, tc.want)
			}
		})
	}
}
