package arraybasics

import "testing"

// Tests for MaxProfit(prices []int) int

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		exp  int
	}{
		{"empty", []int{}, 0},
		{"no_profit", []int{5, 4, 3, 2}, 0},
		{"simple", []int{7, 1, 5, 3, 6, 4}, 5},
		{"peak_late", []int{1, 2, 3, 4, 5}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxProfit(tt.in)
			if got != tt.exp {
				t.Fatalf("MaxProfit(%v) = %d, want %d", tt.in, got, tt.exp)
			}
		})
	}
}
