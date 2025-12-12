
package arraybasics

import "testing"

// Tests for kadane(arr []int) int -> maximum subarray sum

func TestKadane(t *testing.T) {
	tests := []struct{
		name string
		in   []int
		exp  int
	}{
		{"empty", []int{}, 0},
		{"all_negative", []int{-3, -1, -2}, -1},
		{"mixed", []int{-2,1,-3,4,-1,2,1,-5,4}, 6},
		{"all_positive", []int{1,2,3,4}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kadane(tt.in)
			if got != tt.exp {
				t.Fatalf("kadane(%v) = %d, want %d", tt.in, got, tt.exp)
			}
		})
	}
}
