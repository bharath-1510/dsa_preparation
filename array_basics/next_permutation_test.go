
package arraybasics

import (
	"reflect"
	"testing"
)

// Tests for nextPermutation(nums []int) which mutates nums in-place to the
// lexicographically next permutation.
// Remove the build tag when ready to run these against an implementation.

func TestNextPermutation(t *testing.T) {
	tests := []struct{
		name string
		in   []int
		exp  []int
	}{
		{"simple", []int{1,2,3}, []int{1,3,2}},
		{"last_perm", []int{3,2,1}, []int{1,2,3}},
		{"middle", []int{1,3,2}, []int{2,1,3}},
		{"with_duplicates", []int{1,5,1}, []int{5,1,1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]int, len(tt.in))
			copy(arr, tt.in)
			nextPermutation(arr)
			if !reflect.DeepEqual(arr, tt.exp) {
				t.Fatalf("nextPermutation(%v) -> %v, want %v", tt.in, arr, tt.exp)
			}
		})
	}
}
