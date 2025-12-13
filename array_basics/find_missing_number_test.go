package arraybasics

import "testing"

// Tests for FindMissingNumber(nums []int) int

func TestFindMissingNumber(t *testing.T) {
	tests := []struct{
		name string
		in []int
		exp int
	}{
		{"example1", []int{0,1,3}, 2},
		{"example2", []int{3,0,1}, 2},
		{"edge_small", []int{0}, 1},
		{"larger", []int{0,1,2,4,5}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMissingNumber(tt.in)
			if got != tt.exp {
				t.Fatalf("FindMissingNumber(%v) = %d; want %d", tt.in, got, tt.exp)
			}
		})
	}
}
