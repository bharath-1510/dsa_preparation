package binarysearch

import "testing"

func TestAggressiveCows(t *testing.T) {
	tests := []struct {
		name     string
		stalls   []int
		cows     int
		expected int
	}{
		{
			name:     "Base Case - Example 1",
			stalls:   []int{1, 2, 8, 4, 9},
			cows:     3,
			expected: 3,
		},
		{
			name:     "Unsorted Stalls",
			stalls:   []int{10, 1, 2, 7, 5},
			cows:     3,
			expected: 4,
		},
		{
			name:     "Tight Stall Placement",
			stalls:   []int{0, 3, 4, 7, 10, 9},
			cows:     4,
			expected: 3,
		},
		{
			name:     "Maximum Distance with 2 Cows",
			stalls:   []int{5, 4, 3, 2, 1},
			cows:     2,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := AggressiveCows(tt.stalls, tt.cows)
			if actual != tt.expected {
				t.Errorf("AggressiveCows(%v, %d) = %d; want %d", tt.stalls, tt.cows, actual, tt.expected)
			}
		})
	}
}
