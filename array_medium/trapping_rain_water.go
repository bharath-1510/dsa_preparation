package arraymedium

// TrappingRainWater returns the total units of water that can be trapped after raining.
// func TrappingRainWater(arr []int) int {
// 	n := len(arr)
// 	if n == 0 {
// 		return 0
// 	}

// 	leftMax := make([]int, n)
// 	rightMax := make([]int, n)

// 	leftMax[0] = arr[0]
// 	for i := 1; i < n; i++ {
// 		leftMax[i] = max(leftMax[i-1], arr[i])
// 	}

// 	rightMax[n-1] = arr[n-1]
// 	for i := n - 2; i >= 0; i-- {
// 		rightMax[i] = max(rightMax[i+1], arr[i])
// 	}

// 	water := 0
// 	for i := 0; i < n; i++ {
// 		water += min(leftMax[i], rightMax[i]) - arr[i]
// 	}

// 	return water
// }

// TrappingRainWater returns the total units of water that can be trapped after raining.
// Achieved: Time Complexity: O(n), Space Complexity: O(1) — two-pointer approach.
// Best achievable: Time Complexity: O(n), Space Complexity: O(1) — already optimal; alternative O(n) time, O(n) space with precomputed arrays.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TrappingRainWater(arr []int) int {
	left, right := 0, len(arr)-1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		if arr[left] < arr[right] {
			if arr[left] >= leftMax {
				leftMax = arr[left]
			} else {
				water += leftMax - arr[left]
			}
			left++
		} else {
			if arr[right] >= rightMax {
				rightMax = arr[right]
			} else {
				water += rightMax - arr[right]
			}
			right--
		}
	}

	return water
}
