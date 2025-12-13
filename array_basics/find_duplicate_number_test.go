package arraybasics

import "testing"

// Tests for FindDuplicateNumber(nums []int) int

func TestFindDuplicateNumber(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		exp  int
	}{
		{"example1", []int{1, 3, 4, 2, 2}, 2},
		{"example2", []int{3, 1, 3, 4, 2}, 3},
		{"small", []int{2, 2}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindDuplicateNumber(tt.in)
			if got != tt.exp {
				t.Fatalf("FindDuplicateNumber(%v) = %d; want %d", tt.in, got, tt.exp)
			}
		})
	}
}
