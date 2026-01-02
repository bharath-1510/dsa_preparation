package binarysearch

import "testing"

func TestMedianOfTwoSortedArrays(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected float64
	}{
		{
			name:     "Standard Case - Odd Total Length",
			nums1:    []int{1, 3},
			nums2:    []int{2},
			expected: 2.0,
		},
		{
			name:     "Standard Case - Even Total Length",
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: 2.5,
		},
		{
			name:     "One Array Empty",
			nums1:    []int{},
			nums2:    []int{1},
			expected: 1.0,
		},
		{
			name:     "Both Arrays Single Element",
			nums1:    []int{1},
			nums2:    []int{2},
			expected: 1.5,
		},
		{
			name:     "Disjoint Arrays",
			nums1:    []int{1, 2, 3},
			nums2:    []int{10, 20, 30},
			expected: 6.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := MedianOfTwoSortedArrays(tt.nums1, tt.nums2)
			if actual != tt.expected {
				t.Errorf("MedianOfTwoSortedArrays(%v, %v) = %f; want %f", tt.nums1, tt.nums2, actual, tt.expected)
			}
		})
	}
}
