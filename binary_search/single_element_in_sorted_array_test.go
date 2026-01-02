package binarysearch

import "testing"

func TestSingleElementInSortedArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Standard Case - Example 1",
			nums:     []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
			expected: 2,
		},
		{
			name:     "Single Element at the Start",
			nums:     []int{10, 11, 11, 12, 12},
			expected: 10,
		},
		{
			name:     "Single Element at the End",
			nums:     []int{1, 1, 2, 2, 3},
			expected: 3,
		},
		{
			name:     "Array with Only One Element",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "Larger Array - Single Element in Middle",
			nums:     []int{3, 3, 7, 7, 10, 11, 11},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := SingleElementInSortedArray(tt.nums)
			if actual != tt.expected {
				t.Errorf("SingleElementInSortedArray(%v) = %d; want %d", tt.nums, actual, tt.expected)
			}
		})
	}
}
