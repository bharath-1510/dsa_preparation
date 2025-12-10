package arraybasics

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicatesInArray(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		exp  []int
	}{
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
		{"no_duplicates", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"all_duplicates", []int{5, 5, 5, 5}, []int{5}},
		{"mixed_duplicates", []int{1, 1, 2, 2, 3, 3, 3, 4}, []int{1, 2, 3, 4}},
		{"already_unique", []int{0, 2, 2, 3, 5, 5, 7, 9}, []int{0, 2, 3, 5, 7, 9}},
		{"long_sequence", func() []int {
			// create sorted array with some duplicates
			a := make([]int, 0, 200)
			for i := 0; i < 100; i++ {
				a = append(a, i)
				if i%3 == 0 {
					a = append(a, i)
				}
			}
			return a
		}(), func() []int {
			b := make([]int, 0, 100)
			for i := 0; i < 100; i++ {
				b = append(b, i)
			}
			return b
		}()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicatesInArray(tt.in)
			if !reflect.DeepEqual(got, tt.exp) {
				t.Errorf("removeDuplicatesInArray(%v) = %v, want %v", tt.in, got, tt.exp)
			}
		})
	}
}
