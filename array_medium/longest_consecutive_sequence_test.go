package arraymedium

import "testing"

// Assume function signature: func LongestConsecutiveSequence(arr []int) int
func TestLongestConsecutiveSequence(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"example", []int{100, 4, 200, 1, 3, 2, 2, 5}, 5},
		{"single", []int{1}, 1},
		{"empty", []int{}, 0},
		{"duplicates", []int{1, 2, 2, 3}, 3},
		{"no_consecutive", []int{10, 20, 30}, 1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LongestConsecutiveSequence(tc.arr)
			if got != tc.want {
				t.Errorf("LongestConsecutiveSequence(%v) = %d, want %d", tc.arr, got, tc.want)
			}
		})
	}
}
