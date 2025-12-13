package arraybasics

import (
	"reflect"
	"testing"
)

// Tests for MergeIntervals(intervals [][]int) [][]int

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int
		exp  [][]int
	}{
		{"no_overlap", [][]int{{1, 2}, {3, 4}}, [][]int{{1, 2}, {3, 4}}},
		{"overlap", [][]int{{1, 3}, {2, 6}, {8, 10}}, [][]int{{1, 6}, {8, 10}}},
		{"contained", [][]int{{1, 10}, {2, 3}, {4, 8}}, [][]int{{1, 10}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeIntervals(tt.in)
			if !reflect.DeepEqual(got, tt.exp) {
				t.Fatalf("MergeIntervals(%v) = %v, want %v", tt.in, got, tt.exp)
			}
		})
	}
}
