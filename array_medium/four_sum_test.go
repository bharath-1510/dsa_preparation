package arraymedium

import (
	"testing"
	"reflect"
	"sort"
)

// Assume function signature: func FourSum(arr []int, target int) [][]int
func TestFourSum(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   [][]int
	}{
		{"example", []int{1, 0, -1, 0, -2, 2}, 0, [][]int{
			{-2, -1, 1, 2},
			{-2, 0, 0, 2},
			{-1, 0, 0, 1},
		}},
		{"no_quadruplet", []int{1, 2, 3, 4}, 100, [][]int{}},
		{"duplicates", []int{2, 2, 2, 2, 2}, 8, [][]int{{2,2,2,2}}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FourSum(tc.arr, tc.target)
			// Sort inner slices for comparison
			for _, quad := range got {
				sort.Ints(quad)
			}
			for _, quad := range tc.want {
				sort.Ints(quad)
			}
			if !equalQuadruplets(got, tc.want) {
				t.Errorf("FourSum(%v, %d) = %v, want %v", tc.arr, tc.target, got, tc.want)
			}
		})
	}
}

func equalQuadruplets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	// Sort outer slice for comparison
	sort.Slice(a, func(i, j int) bool {
		for k := range a[i] {
			if a[i][k] != a[j][k] {
				return a[i][k] < a[j][k]
			}
		}
		return false
	})
	sort.Slice(b, func(i, j int) bool {
		for k := range b[i] {
			if b[i][k] != b[j][k] {
				return b[i][k] < b[j][k]
			}
		}
		return false
	})
	return reflect.DeepEqual(a, b)
}
