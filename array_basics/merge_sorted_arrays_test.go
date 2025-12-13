package arraybasics

import (
	"reflect"
	"testing"
)

// Tests for MergeSortedArrays(a, b []int) []int

func TestMergeSortedArrays(t *testing.T) {
	tests := []struct{
		name string
		a []int
		b []int
		exp []int
	}{
		{"both_empty", []int{}, []int{}, []int{}},
		{"one_empty", []int{1,2}, []int{}, []int{1,2}},
		{"interleave", []int{1,3,5}, []int{2,4,6}, []int{1,2,3,4,5,6}},
		{"with_duplicates", []int{1,2,2}, []int{2,3}, []int{1,2,2,2,3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSortedArrays(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.exp) {
				t.Fatalf("MergeSortedArrays(%v,%v) = %v, want %v", tt.a, tt.b, got, tt.exp)
			}
		})
	}
}
