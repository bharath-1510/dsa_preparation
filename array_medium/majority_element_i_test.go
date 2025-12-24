package arraymedium

import "testing"

// Assume function signature: func MajorityElementI(arr []int) int
func TestMajorityElementI(t *testing.T) {
	tests := []struct {
		name  string
		arr   []int
		want  int
	}{
		{"example", []int{3, 3, 4, 2, 3}, 3},
		{"all_same", []int{2, 2, 2, 2}, 2},
		{"single", []int{1}, 1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MajorityElementI(tc.arr)
			if got != tc.want {
				t.Errorf("MajorityElementI(%v) = %d, want %d", tc.arr, got, tc.want)
			}
		})
	}
}
