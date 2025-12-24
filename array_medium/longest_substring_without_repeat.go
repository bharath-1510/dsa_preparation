package arraymedium

// LongestSubstringWithoutRepeat returns the length of the longest substring without repeating characters.
// Achieved: Time Complexity: O(n), Space Complexity: O(k) — sliding window with hashmap, k = charset size.
// Best achievable: Time Complexity: O(n), Space Complexity: O(k) — already optimal for this problem.
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
