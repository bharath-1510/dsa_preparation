package arraybasics

import (
	"reflect"
	"testing"
)

// Tests for pascalsTriangle(numRows int) [][]int
// Note: this file is excluded from normal `go test` runs by the build tag at
// the top. Remove the `//go:build ignore` line when implementations exist and
// you want these tests to run.

func TestPascalsTriangle(t *testing.T) {
	tests := []struct{
		name string
		n    int
		exp  [][]int
	}{
		{"zero", 0, [][]int{}},
		{"one", 1, [][]int{{1}}},
		{"three", 3, [][]int{{1},{1,1},{1,2,1}}},
		{"five", 5, [][]int{{1},{1,1},{1,2,1},{1,3,3,1},{1,4,6,4,1}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pascalsTriangle(tt.n)
			if !reflect.DeepEqual(got, tt.exp) {
				t.Fatalf("pascalsTriangle(%d) = %v, want %v", tt.n, got, tt.exp)
			}
		})
	}
}
