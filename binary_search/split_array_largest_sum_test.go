package binarysearch

import "testing"

func TestSplitArrayLargestSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		m        int
		expected int
	}{
		{
			name:     "Base Case - Example 1",
			nums:     []int{7, 2, 5, 10, 8},
			m:        2,
			expected: 18,
		},
		{
			name:     "Split into 5 Subarrays (Each element its own subarray)",
			nums:     []int{1, 2, 3, 4, 5},
			m:        5,
			expected: 5,
		},
		{
			name:     "Split into 1 Subarray (Sum of all elements)",
			nums:     []int{1, 2, 3, 4, 5},
			m:        1,
			expected: 15,
		},
		{
			name:     "Uniform values",
			nums:     []int{10, 10, 10, 10},
			m:        2,
			expected: 20,
		},
		{
			name:     "Single element",
			nums:     []int{5},
			m:        1,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := SplitArrayLargestSum(tt.nums, tt.m)
			if actual != tt.expected {
				t.Errorf("SplitArrayLargestSum(%v, %d) = %d; want %d", tt.nums, tt.m, actual, tt.expected)
			}
		})
	}
}
