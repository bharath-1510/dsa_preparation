package strings_fundamentals

// LongestPalindromicSubstring finds the longest palindromic substring in s.
// Time: O(n^2), Space: O(1) — achieved
// Best possible: O(n) time (Manacher's algorithm), O(n) space
func LongestPalindromicSubstring(s string) string {
	if len(s) == 0 {
		return ""
	}

	r := []rune(s)
	start, maxLen := 0, 1

	for i := 0; i < len(r); i++ {
		l1, len1 := expand(r, i, i)
		if len1 > maxLen {
			start = l1
			maxLen = len1
		}

		l2, len2 := expand(r, i, i+1)
		if len2 > maxLen {
			start = l2
			maxLen = len2
		}
	}

	return string(r[start : start+maxLen])
}

func expand(r []rune, left, right int) (int, int) {
	for left >= 0 && right < len(r) && r[left] == r[right] {
		left--
		right++
	}

	start := left + 1
	length := right - left - 1

	return start, length
}
