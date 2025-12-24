package arraymedium

import "testing"
import "sort"

// Assume function signature: func MajorityElementII(arr []int) []int
func TestMajorityElementII(t *testing.T) {
	tests := []struct {
		name  string
		arr   []int
		want  []int
	}{
		{"example", []int{3, 2, 3, 2, 2, 3, 3}, []int{2, 3}},
		{"none", []int{1, 2, 3}, []int{}},
		{"all_same", []int{2, 2, 2, 2}, []int{2}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MajorityElementII(tc.arr)
			sort.Ints(got)
			sort.Ints(tc.want)
			if len(got) != len(tc.want) {
				t.Errorf("MajorityElementII(%v) = %v, want %v", tc.arr, got, tc.want)
				return
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Errorf("MajorityElementII(%v) = %v, want %v", tc.arr, got, tc.want)
					return
				}
			}
		})
	}
}
