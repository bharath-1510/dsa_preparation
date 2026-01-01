package binarysearch

import "testing"

func TestPaintersPartition(t *testing.T) {
	tests := []struct {
		name     string
		boards   []int
		painters int
		expected int
	}{
		{
			name:     "Base Case - Example 1",
			boards:   []int{10, 20, 30, 40, 50},
			painters: 2,
			expected: 90,
		},
		{
			name:     "Uniform Board Lengths",
			boards:   []int{10, 10, 10, 10},
			painters: 2,
			expected: 20,
		},
		{
			name:     "Each Painter Gets One Board",
			boards:   []int{5, 5, 5, 5},
			painters: 4,
			expected: 5,
		},
		{
			name:     "Multiple Boards per Painter",
			boards:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			painters: 3,
			expected: 17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := PaintersPartition(tt.boards, tt.painters)
			if actual != tt.expected {
				t.Errorf("PaintersPartition(%v, %d) = %d; want %d", tt.boards, tt.painters, actual, tt.expected)
			}
		})
	}
}
