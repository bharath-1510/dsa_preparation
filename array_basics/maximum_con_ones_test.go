package arraybasics

import "testing"

func TestMaximumConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		exp  int
	}{
		{"empty", []int{}, 0},
		{"all_zero", []int{0, 0, 0}, 0},
		{"all_one", []int{1, 1, 1, 1}, 4},
		{"mixed", []int{1, 1, 0, 1, 1, 1, 0, 1}, 3},
		{"single_one", []int{0, 0, 1, 0}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumConsecutiveOnes(tt.in)
			if got != tt.exp {
				t.Fatalf("maximumConsecutiveOnes(%v) = %d, want %d", tt.in, got, tt.exp)
			}
		})
	}
}
