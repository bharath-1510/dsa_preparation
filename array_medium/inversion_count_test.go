package arraymedium

import "testing"

// Assume function signature: func InversionCount(arr []int) int
func TestInversionCount(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"example", []int{2, 4, 1, 3, 5}, 3},
		{"sorted", []int{1, 2, 3, 4, 5}, 0},
		{"reverse", []int{5, 4, 3, 2, 1}, 10},
		{"single", []int{1}, 0},
		{"empty", []int{}, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := InversionCount(tc.arr)
			if got != tc.want {
				t.Errorf("InversionCount(%v) = %d, want %d", tc.arr, got, tc.want)
			}
		})
	}
}
