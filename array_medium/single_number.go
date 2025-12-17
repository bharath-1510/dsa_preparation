package arraymedium

// func SingleNumber(arr []int) int {
// 	out := 0
// 	res := make(map[int]bool)
// 	for _, val := range arr {
// 		if res[val] {
// 			res[val] = false
// 		} else {
// 			res[val] = true
// 		}
// 	}
// 	for k, v := range res {
// 		if v {
// 			return k
// 		}
// 	}
// 	return out
// }

func SingleNumber(arr []int) int {
	res := 0
	for _, v := range arr {
		res ^= v
	}
	return res
}