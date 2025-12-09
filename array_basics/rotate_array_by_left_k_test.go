package arraybasics

import (
    "reflect"
    "testing"
)

func TestRotateArrayByLeftK(t *testing.T) {
    tests := []struct {
        name string
        arr  []int
        k    int
        want []int
    }{
        {name: "empty slice", arr: []int{}, k: 3, want: []int{}},
        {name: "k zero", arr: []int{1, 2, 3, 4}, k: 0, want: []int{1, 2, 3, 4}},
        {name: "k one", arr: []int{1, 2, 3, 4}, k: 1, want: []int{2, 3, 4, 1}},
        {name: "k equals len", arr: []int{1, 2, 3, 4}, k: 4, want: []int{1, 2, 3, 4}},
        {name: "k greater than len", arr: []int{1, 2, 3, 4}, k: 5, want: []int{2, 3, 4, 1}},
        {name: "duplicates", arr: []int{2, 2, 3, 2}, k: 2, want: []int{3, 2, 2, 2}},
        {name: "single element", arr: []int{42}, k: 10, want: []int{42}},
        {name: "large k", arr: []int{1, 2, 3, 4, 5}, k: 12, want: []int{3, 4, 5, 1, 2}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := rotate_array_by_left_k(tt.arr, tt.k)
            if !reflect.DeepEqual(got, tt.want) {
                t.Fatalf("rotate_array_by_left_k(%v, %d) = %v; want %v", tt.arr, tt.k, got, tt.want)
            }
        })
    }
}
