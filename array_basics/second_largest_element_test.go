package arraybasics

import "testing"

func TestSecondLargestElement(t *testing.T) {
    tests := []struct {
        name string
        arr  []int
        want int
    }{
        {name: "normal", arr: []int{1, 3, 2, 5, 4}, want: 4},
        {name: "duplicates with distinct second", arr: []int{5, 5, 4, 3}, want: 4},
        {name: "all duplicates", arr: []int{2, 2, 2}, want: 0},
        {name: "single element", arr: []int{42}, want: 0},
        {name: "empty slice", arr: []int{}, want: 0},
        {name: "negative numbers", arr: []int{-10, -3, -7, -1}, want: -3},
        {name: "mixed", arr: []int{0, -1, 100, 50}, want: 50},
        {name: "second is negative", arr: []int{10, -5, -1, 10}, want: -1},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := second_largest_element(tt.arr)
            if got != tt.want {
                t.Fatalf("second_largest_element(%v) = %d; want %d", tt.arr, got, tt.want)
            }
        })
    }
}
