package arraymedium

// func TwoSum(arr []int, target int) (int, int) {
// 	for i := range arr {
// 		for j := i + 1; j < len(arr); j++ {
// 			if arr[i]+arr[j] == target {
// 				return i, j
// 			}
// 		}
// 	}
// 	return -1, -1
// }

func TwoSum(arr []int, target int) (int, int) {
	indexMap := make(map[int]int)
	for i, val := range arr {
		need := target - val
		if idx, ok := indexMap[need]; ok {
			return idx, i
		}
		indexMap[val] = i
	}
	return -1, -1
}