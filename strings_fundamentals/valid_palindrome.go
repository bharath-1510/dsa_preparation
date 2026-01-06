package strings_fundamentals

import "strings"

// ValidPalindrome checks if a string is a palindrome (alphanumeric, ignore case).
func ValidPalindrome(s string) bool {
	lowercaseStr := strings.ToLower(s)
	runes := []rune(lowercaseStr)
	left, right := 0, len(runes)-1
	for left <= right {
		leftChar, rightChar := runes[left], runes[right]
		if !checkAlphaNumberic(leftChar) {
			left++
			continue
		}
		if !checkAlphaNumberic(rightChar) {
			right--
			continue
		}
		if leftChar != rightChar {
			return false
		}
		left++
		right--

	}
	return true
}

func checkAlphaNumberic(ch rune) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
		return true
	}
	return false
}
