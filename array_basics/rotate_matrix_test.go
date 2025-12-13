package arraybasics

import (
	"reflect"
	"testing"
)

// Tests for RotateMatrix(mat [][]int) [][]int (NxN rotation 90° clockwise)

func TestRotateMatrix(t *testing.T) {
	tests := []struct{
		name string
		in [][]int
		exp [][]int
	}{
		{"2x2", [][]int{{1,2},{3,4}}, [][]int{{3,1},{4,2}}},
		{"3x3", [][]int{{1,2,3},{4,5,6},{7,8,9}}, [][]int{{7,4,1},{8,5,2},{9,6,3}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RotateMatrix(tt.in)
			if !reflect.DeepEqual(got, tt.exp) {
				t.Fatalf("RotateMatrix(%v) = %v, want %v", tt.in, got, tt.exp)
			}
		})
	}
}
