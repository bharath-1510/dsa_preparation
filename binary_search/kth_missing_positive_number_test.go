package binarysearch

import "testing"

func TestKthMissingPositiveNumber(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		k        int
		expected int
	}{
		{
			name:     "Base Case - Example 1",
			arr:      []int{2, 3, 4, 7, 11},
			k:        5,
			expected: 9,
		},
		{
			name:     "Missing numbers before the first element",
			arr:      []int{5, 6, 7},
			k:        2,
			expected: 2,
		},
		{
			name:     "Kth missing number is very large (beyond array)",
			arr:      []int{1, 2, 3},
			k:        2,
			expected: 5,
		},
		{
			name:     "Array starts with 1",
			arr:      []int{1, 3, 5},
			k:        2,
			expected: 4,
		},
		{
			name:     "Single element array",
			arr:      []int{10},
			k:        5,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := KthMissingPositiveNumber(tt.arr, tt.k)
			if actual != tt.expected {
				t.Errorf("KthMissingPositiveNumber(%v, %d) = %d; want %d", tt.arr, tt.k, actual, tt.expected)
			}
		})
	}
}
