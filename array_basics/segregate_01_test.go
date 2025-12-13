package arraybasics

import (
	"reflect"
	"testing"
)

// Tests for Segregate01(arr []int) []int

func TestSegregate01(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		exp  []int
	}{
		{"empty", []int{}, []int{}},
		{"already", []int{0, 0, 1, 1}, []int{0, 0, 1, 1}},
		{"mixed", []int{0, 1, 0, 1, 1, 0}, []int{0, 0, 0, 1, 1, 1}},
		{"all_zero", []int{0, 0, 0}, []int{0, 0, 0}},
		{"all_one", []int{1, 1}, []int{1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Segregate01(tt.in)
			if !reflect.DeepEqual(got, tt.exp) {
				t.Fatalf("Segregate01(%v) = %v, want %v", tt.in, got, tt.exp)
			}
		})
	}
}
