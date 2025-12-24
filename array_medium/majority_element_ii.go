package arraymedium

// MajorityElementII returns all elements that appear more than n/3 times.
// func MajorityElementII(arr []int) []int {
// 	res := make([]int, 0)
// 	freq := make(map[int]int, 0)
// 	for _, val := range arr {
// 		freq[val]++
// 	}
// 	target := len(arr)/3 + 1
// 	for k, v := range freq {
// 		if v >= target {
// 			res = append(res, k)
// 		}
// 	}
// 	return res
// }

func MajorityElementII(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}
	var cand1, cand2 int
	count1, count2 := 0, 0

	for _, v := range arr {
		if count1 > 0 && v == cand1 {
			count1++
		} else if count2 > 0 && v == cand2 {
			count2++
		} else if count1 == 0 {
			cand1 = v
			count1 = 1
		} else if count2 == 0 {
			cand2 = v
			count2 = 1
		} else {
			count1--
			count2--
		}
	}

	count1, count2 = 0, 0
	for _, v := range arr {
		if v == cand1 {
			count1++
		} else if v == cand2 {
			count2++
		}
	}

	result := make([]int, 0)
	if count1 > n/3 {
		result = append(result, cand1)
	}
	if cand2 != cand1 && count2 > n/3 {
		result = append(result, cand2)
	}

	return result
}
