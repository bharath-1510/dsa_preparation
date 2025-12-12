package arraybasics

import (
	"reflect"
	"testing"
)

// Tests for sortColors(nums []int) which sorts 0/1/2 in-place.

func TestSortColors(t *testing.T) {
	tests := []struct{
		name string
		in   []int
		exp  []int
	}{
		{"empty", []int{}, []int{}},
		{"already_sorted", []int{0,0,1,1,2,2}, []int{0,0,1,1,2,2}},
		{"mixed", []int{2,0,2,1,1,0}, []int{0,0,1,1,2,2}},
		{"single_type", []int{1,1,1}, []int{1,1,1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]int, len(tt.in))
			copy(arr, tt.in)
			sortColors(arr)
			if !reflect.DeepEqual(arr, tt.exp) {
				t.Fatalf("sortColors(%v) -> %v, want %v", tt.in, arr, tt.exp)
			}
		})
	}
}
