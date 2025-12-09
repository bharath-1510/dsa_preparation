package arraybasics

import "testing"

func TestLargestElement(t *testing.T) {
    tests := []struct {
        name string
        arr  []int
        want int
    }{
        {name: "normal", arr: []int{1, 3, 2, 5, 4}, want: 5},
        {name: "negative numbers", arr: []int{-10, -3, -7, -1}, want: -1},
        {name: "single element", arr: []int{42}, want: 42},
        {name: "duplicates", arr: []int{2, 2, 2, 2}, want: 2},
        {name: "empty slice", arr: []int{}, want: 0},
        {name: "mixed", arr: []int{0, -1, 100, 50}, want: 100},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := largestElement(tt.arr)
            if got != tt.want {
                t.Fatalf("largestElement(%v) = %d; want %d", tt.arr, got, tt.want)
            }
        })
    }
}
