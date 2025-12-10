package arraybasics

import "testing"

func TestIsSorted(t *testing.T) {
    tests := []struct{
        name string
        arr  []int
        want bool
    }{
        {name: "empty slice", arr: []int{}, want: true},
        {name: "single element", arr: []int{42}, want: true},
        {name: "already sorted", arr: []int{1,2,2,3,5}, want: true},
        {name: "strictly increasing", arr: []int{1,2,3,4,5}, want: true},
        {name: "unsorted at end", arr: []int{1,2,5,4}, want: false},
        {name: "unsorted in middle", arr: []int{3,2,1}, want: false},
        {name: "all equal", arr: []int{2,2,2,2}, want: true},
        {name: "negative numbers sorted", arr: []int{-5, -3, -3, 0}, want: true},
        {name: "negative unsorted", arr: []int{-1, -2, 0}, want: false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T){
            got := isSorted(tt.arr)
            if got != tt.want {
                t.Fatalf("isSorted(%v) = %v; want %v", tt.arr, got, tt.want)
            }
        })
    }
}
