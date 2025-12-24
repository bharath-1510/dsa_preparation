package arraymedium

import "testing"

// Assume function signature: func NumberOfSubarraysWithSumK(arr []int, k int) int
func TestNumberOfSubarraysWithSumK(t *testing.T) {
	tests := []struct {
		name  string
		arr   []int
		k     int
		want  int
	}{
		{"example", []int{1, 1, 1, 2, 1}, 3, 3},
		{"none", []int{1, 2, 3}, 100, 0},
		{"negatives", []int{1, -1, 1, 1}, 1, 5},
		{"empty", []int{}, 10, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := NumberOfSubarraysWithSumK(tc.arr, tc.k)
			if got != tc.want {
				t.Errorf("NumberOfSubarraysWithSumK(%v, %d) = %d, want %d", tc.arr, tc.k, got, tc.want)
			}
		})
	}
}
