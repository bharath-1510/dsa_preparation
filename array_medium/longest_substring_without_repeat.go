package arraymedium

func LongestSubstringWithoutRepeat(s string) int {
	seen := make(map[rune]int)
	maxLen := 0
	l := 0

	for r, ch := range []rune(s) {
		if idx, ok := seen[ch]; ok && idx >= l {
			l = idx + 1
		}
		seen[ch] = r
		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
	}

	return maxLen
}
