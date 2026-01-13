package strings_fundamentals

// CountPalindromicSubstrings counts the total number of palindromic substrings in s.
// Time: O(n^2), Space: O(1) — achieved
// Best possible: O(n^2) time, O(1) space (optimal for this problem)
func CountPalindromicSubstrings(s string) int {
	runes := []rune(s)
	n := len(runes)
	count := 0

	for i := 0; i < n; i++ {
		count += expandAroundCenter(runes, i, i)
		count += expandAroundCenter(runes, i, i+1)
	}

	return count
}

func expandAroundCenter(r []rune, left, right int) int {
	count := 0

	for left >= 0 && right < len(r) && r[left] == r[right] {
		count++
		left--
		right++
	}

	return count
}
