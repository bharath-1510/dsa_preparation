package binarysearch

import "testing"

func TestFloorCeil(t *testing.T) {
	tests := []struct {
		name          string
		nums          []int
		x             int
		expectedFloor int
		expectedCeil  int
	}{
		{
			name:          "Internal Value - Example 1",
			nums:          []int{1, 2, 8, 10, 10, 12, 19},
			x:             5,
			expectedFloor: 2,
			expectedCeil:  8,
		},
		{
			name:          "Value Exists in Array",
			nums:          []int{1, 2, 8, 10, 10, 12, 19},
			x:             8,
			expectedFloor: 8,
			expectedCeil:  8,
		},
		{
			name:          "Value Smaller than Smallest",
			nums:          []int{5, 10, 15},
			x:             2,
			expectedFloor: -1,
			expectedCeil:  5,
		},
		{
			name:          "Value Larger than Largest",
			nums:          []int{5, 10, 15},
			x:             20,
			expectedFloor: 15,
			expectedCeil:  -1,
		},
		{
			name:          "Single Element - Smaller",
			nums:          []int{10},
			x:             5,
			expectedFloor: -1,
			expectedCeil:  10,
		},
		{
			name:          "Single Element - Larger",
			nums:          []int{10},
			x:             15,
			expectedFloor: 10,
			expectedCeil:  -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			floor, ceil := FloorCeil(tt.nums, tt.x)
			if floor != tt.expectedFloor || ceil != tt.expectedCeil {
				t.Errorf("FloorCeil(%v, %d) = (Floor:%d, Ceil:%d); want (Floor:%d, Ceil:%d)",
					tt.nums, tt.x, floor, ceil, tt.expectedFloor, tt.expectedCeil)
			}
		})
	}
}
