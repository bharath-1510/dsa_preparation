package strings_fundamentals

import "strings"

// ReverseWordsInString reverses the order of words in a string.
func ReverseWordsInString(s string) string {
	words := strings.Split(s, " ")
	left, right := 0, len(words)-1
	for left <= right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}
	return strings.Join(words, " ")
}
