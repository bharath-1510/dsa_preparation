package arraymedium

// func LongestConsecutiveSequence(arr []int) int {
// 	if len(arr) < 2 {
// 		return len(arr)
// 	}
// 	maxCount := 0
// 	res := make([]int, 0)
// 	res = append(res, arr...)
// 	sort.Ints(res)
// 	count := 1
// 	for i := 0; i < len(res)-1; i++ {
// 		if res[i] == res[i+1] {
// 			continue
// 		} else {
// 			if res[i]+1 == res[i+1] {
// 				count++
// 			} else {
// 				if maxCount < count {
// 					maxCount = count
// 				}
// 				count = 1
// 			}
// 		}
// 	}
// 	if maxCount < count {
// 		maxCount = count
// 	}
// 	return maxCount
// }

func LongestConsecutiveSequence(arr []int) int {
	set := make(map[int]bool)
	for _, v := range arr {
		set[v] = true
	}

	maxLen := 0
	for _, v := range arr {
		if !set[v-1] {
			current := v
			count := 1
			for set[current+1] {
				current++
				count++
			}
			if count > maxLen {
				maxLen = count
			}
		}
	}
	return maxLen
}
