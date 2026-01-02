package binarysearch

import "testing"

func TestKthElementOfTwoSortedArrays(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		k        int
		expected int
	}{
		{
			name:     "Standard Case - Example 1",
			nums1:    []int{2, 3, 6, 7, 9},
			nums2:    []int{1, 4, 8, 10},
			k:        5,
			expected: 6,
		},
		{
			name:     "k is 1 (Smallest Element)",
			nums1:    []int{2, 3, 6, 7, 9},
			nums2:    []int{1, 4, 8, 10},
			k:        1,
			expected: 1,
		},
		{
			name:     "k is total length (Largest Element)",
			nums1:    []int{2, 3},
			nums2:    []int{1, 4},
			k:        4,
			expected: 4,
		},
		{
			name:     "One Array empty",
			nums1:    []int{},
			nums2:    []int{10, 20, 30},
			k:        2,
			expected: 20,
		},
		{
			name:     "kth element at the start of nums2",
			nums1:    []int{1, 5, 10},
			nums2:    []int{2, 3, 4},
			k:        2,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := KthElementOfTwoSortedArrays(tt.nums1, tt.nums2, tt.k)
			if actual != tt.expected {
				t.Errorf("KthElementOfTwoSortedArrays(%v, %v, %d) = %d; want %d", tt.nums1, tt.nums2, tt.k, actual, tt.expected)
			}
		})
	}
}
