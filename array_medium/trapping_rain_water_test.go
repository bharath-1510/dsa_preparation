package arraymedium

import "testing"

// Assume function signature: func TrappingRainWater(arr []int) int
func TestTrappingRainWater(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"example", []int{1, 0, 0, 0, 2, 0, 1}, 4},
		{"no_trap", []int{0, 1, 2, 3}, 0},
		{"all_zero", []int{0, 0, 0, 0}, 0},
		{"empty", []int{}, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TrappingRainWater(tc.arr)
			if got != tc.want {
				t.Errorf("TrappingRainWater(%v) = %d, want %d", tc.arr, got, tc.want)
			}
		})
	}
}
