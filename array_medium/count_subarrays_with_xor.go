package arraymedium

// func CountSubarraysWithXOR(arr []int, x int) int {
// 	ans := 0
// 	n := len(arr)

// 	for i := range n {
// 		xor := 0
// 		for j := i; j < n; j++ {
// 			xor ^= arr[j]
// 			if xor == x {
// 				ans++
// 			}
// 		}
// 	}
// 	return ans
// }

func CountSubarraysWithXOR(arr []int, x int) int {
	count := 0
	prefixXor := 0

	freq := make(map[int]int)
	freq[0] = 1

	for _, val := range arr {
		prefixXor ^= val
		if c, ok := freq[prefixXor^x]; ok {
			count += c
		}
		freq[prefixXor]++
	}

	return count
}
