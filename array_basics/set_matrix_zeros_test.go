package arraybasics

import (
	"reflect"
	"testing"
)

func TestSetMatrixZero(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int
		exp  [][]int
	}{
		{
			"no_zero",
			[][]int{{1, 2}, {3, 4}},
			[][]int{{1, 2}, {3, 4}},
		},
		{
			"single_zero",
			[][]int{{1, 0, 3}, {4, 5, 6}},
			[][]int{{0, 0, 0}, {4, 0, 6}},
		},
		{
			"multiple_zeros",
			[][]int{{0, 2, 3}, {4, 5, 0}, {7, 8, 9}},
			[][]int{{0, 0, 0}, {0, 0, 0}, {0, 8, 0}},
		},
		{
			"all_zero",
			[][]int{{0, 0}, {0, 0}},
			[][]int{{0, 0}, {0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := setMatrixZero(tt.in)
			if !reflect.DeepEqual(got, tt.exp) {
				t.Fatalf("setMatrixZero(%v) = %v, want %v", tt.in, got, tt.exp)
			}
		})
	}
}
