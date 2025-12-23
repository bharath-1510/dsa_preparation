package arraymedium

import "testing"

// Assume function signature: func CountSpecialTriplets(arr []int) int
func TestCountSpecialTriplets(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"example", []int{1, 2, 3, 1}, 4},
		{"example1", []int{1, 2, 3}, 2},
		{"all_zero", []int{0, 0, 0}, 4},
		{"empty", []int{}, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CountSpecialTriplets(tc.arr)
			if got != tc.want {
				t.Errorf("CountSpecialTriplets(%v) = %d, want %d", tc.arr, got, tc.want)
			}
		})
	}
}
