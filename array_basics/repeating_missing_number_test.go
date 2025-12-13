package arraybasics

import "testing"

// Tests for FindRepeatingAndMissing(nums []int) (repeated int, missing int)

func TestFindRepeatingAndMissing(t *testing.T) {
	tests := []struct{
		name string
		in []int
		rep int
		mis int
	}{
		{"example1", []int{3,1,2,5,3}, 3, 4},
		{"small", []int{2,2}, 2, 1},
		{"another", []int{4,3,6,2,1,1}, 1, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep, mis := FindRepeatingAndMissing(tt.in)
			if rep != tt.rep || mis != tt.mis {
				t.Fatalf("FindRepeatingAndMissing(%v) = (%d,%d); want (%d,%d)", tt.in, rep, mis, tt.rep, tt.mis)
			}
		})
	}
}
