package arraybasics

import (
	"reflect"
	"testing"
)

func TestMoveZeroEnd(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		exp  []int
	}{
		{"empty", []int{}, []int{}},
		{"no_zero", []int{1, 2, 3}, []int{1, 2, 3}},
		{"all_zero", []int{0, 0, 0}, []int{0, 0, 0}},
		{"mixed", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"trailing_zero", []int{1, 2, 0, 0}, []int{1, 2, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := moveZeroEnd(tt.in)
			if !reflect.DeepEqual(got, tt.exp) {
				t.Fatalf("moveZeroEnd(%v) = %v, want %v", tt.in, got, tt.exp)
			}
		})
	}
}
