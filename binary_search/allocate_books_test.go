package binarysearch

import "testing"

func TestAllocateBooks(t *testing.T) {
	tests := []struct {
		name     string
		pages    []int
		students int
		expected int
	}{
		{
			name:     "Base Case - Example 1",
			pages:    []int{12, 34, 67, 90},
			students: 2,
			expected: 113,
		},
		{
			name:     "Each Student Gets One Book",
			pages:    []int{5, 17, 100, 11},
			students: 4,
			expected: 100,
		},
		{
			name:     "Contiguous Allocation Optimization",
			pages:    []int{25, 46, 28, 49, 24},
			students: 4,
			expected: 71,
		},
		{
			name:     "More Students than Books",
			pages:    []int{12, 34, 67, 90},
			students: 5,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := AllocateBooks(tt.pages, tt.students)
			if actual != tt.expected {
				t.Errorf("AllocateBooks(%v, %d) = %d; want %d", tt.pages, tt.students, actual, tt.expected)
			}
		})
	}
}
