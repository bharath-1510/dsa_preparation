package arraymedium

import "sort"

// func FourSum(arr []int, target int) [][]int {
// 	res := make([][]int, 0)
// 	seen := make(map[[4]int]bool)
// 	n := len(arr)
// 	for i := range n {
// 		for j := i + 1; j < n; j++ {
// 			for k := j + 1; k < n; k++ {
// 				for l := k + 1; l < n; l++ {
// 					if arr[i]+arr[j]+arr[k]+arr[l] == target {
// 						quad := []int{arr[i], arr[j], arr[k], arr[l]}

// 						for x := 0; x < 4; x++ {
// 							for y := x + 1; y < 4; y++ {
// 								if quad[x] > quad[y] {
// 									quad[x], quad[y] = quad[y], quad[x]
// 								}
// 							}
// 						}

// 						key := [4]int{quad[0], quad[1], quad[2], quad[3]}
// 						if !seen[key] {
// 							seen[key] = true
// 							res = append(res, quad)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return res
// }

func FourSum(arr []int, target int) [][]int {
	sort.Ints(arr)
	n := len(arr)
	res := [][]int{}

	for i := 0; i < n-3; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}

			left, right := j+1, n-1

			for left < right {
				sum := arr[i] + arr[j] + arr[left] + arr[right]

				if sum == target {
					res = append(res, []int{
						arr[i], arr[j], arr[left], arr[right],
					})

					left++
					right--

					for left < right && arr[left] == arr[left-1] {
						left++
					}
					for left < right && arr[right] == arr[right+1] {
						right--
					}

				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return res
}
