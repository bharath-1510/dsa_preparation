package arraybasics

// func MaxProfit(arr []int) int {
// 	maxProfit := 0
// 	n := len(arr)
// 	for i := 0; i < n-1; i++ {
// 		for j := i + 1; j < n; j++ {
// 			profit := arr[j] - arr[i]
// 			if maxProfit < profit {
// 				maxProfit = profit
// 			}
// 		}
// 	}
// 	return maxProfit
// }

// MaxProfit computes the max profit from a single buy/sell.
// Time Complexity: O(n) — single pass tracking min price and max profit.
// Space Complexity: O(1) — constant extra space.
// Best achievable: Time = O(n), Space = O(1) — already optimal for single-transaction profit.
func MaxProfit(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	minPrice := arr[0]
	maxProfit := 0

	for i := 1; i < len(arr); i++ {
		profit := arr[i] - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
		if arr[i] < minPrice {
			minPrice = arr[i]
		}
	}

	return maxProfit
}
