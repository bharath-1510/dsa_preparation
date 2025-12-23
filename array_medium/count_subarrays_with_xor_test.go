package arraymedium

import "testing"

// Assume function signature: func CountSubarraysWithXOR(arr []int, x int) int
func TestCountSubarraysWithXOR(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		x    int
		want int
	}{
		{"example", []int{4, 2, 2, 6}, 6, 3},
		{"none", []int{1, 2, 3}, 7, 0},
		{"all_zero", []int{0, 0, 0}, 0, 6},
		{"empty", []int{}, 0, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CountSubarraysWithXOR(tc.arr, tc.x)
			if got != tc.want {
				t.Errorf("CountSubarraysWithXOR(%v, %d) = %d, want %d", tc.arr, tc.x, got, tc.want)
			}
		})
	}
}
