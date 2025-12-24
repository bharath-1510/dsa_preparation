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

// SingleNumber returns the element that appears once when all others appear twice.
// Achieved: Time Complexity: O(n), Space Complexity: O(1) — uses XOR.
// Best achievable: Time Complexity: O(n), Space Complexity: O(1) — already optimal.
func SingleNumber(arr []int) int {
	res := 0
	for _, v := range arr {
		res ^= v
	}
	return res
}
