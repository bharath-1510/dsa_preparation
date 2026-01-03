package binarysearch

import "testing"

func TestKokoEatingBananas(t *testing.T) {
	tests := []struct {
		name     string
		piles    []int
		h        int
		expected int
	}{
		{
			name:     "Standard Case - Example 1",
			piles:    []int{3, 6, 7, 11},
			h:        8,
			expected: 4,
		},
		{
			name:     "Eating slowly - One pile per hour",
			piles:    []int{30, 11, 23, 4, 20},
			h:        5,
			expected: 30,
		},
		{
			name:     "More hours than piles",
			piles:    []int{30, 11, 23, 4, 20},
			h:        6,
			expected: 23,
		},
		{
			name:     "Large number of hours",
			piles:    []int{100, 100, 100},
			h:        1000,
			expected: 1,
		},
		{
			name:     "Single pile",
			piles:    []int{10},
			h:        2,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := KokoEatingBananas(tt.piles, tt.h)
			if actual != tt.expected {
				t.Errorf("KokoEatingBananas(%v, %d) = %d; want %d", tt.piles, tt.h, actual, tt.expected)
			}
		})
	}
}
