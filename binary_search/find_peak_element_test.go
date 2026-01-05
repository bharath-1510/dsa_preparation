package binarysearch

import "testing"

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int // possible valid indices
	}{
		{
			name:     "Base Case - Example 1",
			nums:     []int{1, 2, 1, 3, 5, 6, 4},
			expected: []int{1, 5},
		},
		{
			name:     "Strictly Increasing",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{4},
		},
		{
			name:     "Strictly Decreasing",
			nums:     []int{5, 4, 3, 2, 1},
			expected: []int{0},
		},
		{
			name:     "Single Element",
			nums:     []int{1},
			expected: []int{0},
		},
		{
			name:     "Peak in middle",
			nums:     []int{1, 3, 2},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := FindPeakElement(tt.nums)
			found := false
			for _, exp := range tt.expected {
				if actual == exp {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("FindPeakElement(%v) = %d; want one of %v", tt.nums, actual, tt.expected)
			}
		})
	}
}
