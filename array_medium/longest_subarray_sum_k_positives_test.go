package arraymedium

import "testing"

// Assume function signature: func LongestSubarraySumKPositives(arr []int, k int) int
func TestLongestSubarraySumKPositives(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		k    int
		want int
	}{
		{"example", []int{1, 2, 3, 7, 5, 2}, 12, 3},
		{"none", []int{1, 2, 3}, 100, 0},
		{"single", []int{5}, 5, 1},
		{"empty", []int{}, 10, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LongestSubarraySumKPositives(tc.arr, tc.k)
			if got != tc.want {
				t.Errorf("LongestSubarraySumKPositives(%v, %d) = %d, want %d", tc.arr, tc.k, got, tc.want)
			}
		})
	}
}
