package binarysearch

import "testing"

func TestTimeToEatBananas(t *testing.T) {
	tests := []struct {
		name     string
		piles    []int
		h        int
		expected int
	}{
		{
			name:     "Standard Case - Example 1",
			piles:    []int{30, 11, 23, 4, 20},
			h:        5,
			expected: 30,
		},
		{
			name:     "Small input - Fewer hours",
			piles:    []int{3, 6, 7, 11},
			h:        8,
			expected: 4,
		},
		{
			name:     "Single pile - Exact speed",
			piles:    []int{10},
			h:        10,
			expected: 1,
		},
		{
			name:     "Single pile - Half speed",
			piles:    []int{10},
			h:        1,
			expected: 10,
		},
		{
			name:     "All piles same size",
			piles:    []int{10, 10, 10},
			h:        6,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := TimeToEatBananas(tt.piles, tt.h)
			if actual != tt.expected {
				t.Errorf("TimeToEatBananas(%v, %d) = %d; want %d", tt.piles, tt.h, actual, tt.expected)
			}
		})
	}
}
