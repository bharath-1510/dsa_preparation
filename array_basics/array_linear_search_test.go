package arraybasics

import "testing"

func TestArrayLinearSearch(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		exp    int
	}{
		{"empty", []int{}, 1, -1},
		{"single_found", []int{5}, 5, 0},
		{"single_not_found", []int{5}, 3, -1},
		{"multiple_found", []int{1, 2, 3, 4}, 3, 2},
		{"multiple_not_found", []int{1, 2, 3, 4}, 7, -1},
		{"duplicates_returns_first", []int{2, 2, 3, 2}, 2, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := arrayLinearSearch(tt.arr, tt.target)
			if got != tt.exp {
				t.Fatalf("arrayLinearSearch(%v, %d) = %d; want %d", tt.arr, tt.target, got, tt.exp)
			}
		})
	}
}

