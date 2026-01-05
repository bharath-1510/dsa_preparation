# Strings Fundamentals — Problem Statements & Examples

Below are concise problem statements, input/output formats and one example for each problem in this folder.

## Reverse Words in String
Problem: Reverse the order of words in a given string.
Input: A string containing words separated by spaces.
Output: A string with words in reverse order.
Example:
Input: "the sky is blue"
Output: "blue is sky the"

## Longest Palindromic Substring
Problem: Find the longest substring that is a palindrome.
Input: A single string.
Output: Longest palindromic substring.
Example:
Input: "babad"
Output: "bab" (or "aba")

## Count Palindromic Substrings
Problem: Count the total number of palindromic substrings in a string.
Input: A single string.
Output: Total count of palindromic substrings.
Example:
Input: "aaa"
Output: 6

## Longest Common Prefix
Problem: Find the longest common prefix string amongst an array of strings.
Input: An array/slice of strings.
Output: The common prefix string.
Example:
Input: ["flower","flow","flight"]
Output: "fl"

## Valid Anagram
Problem: Determine if two strings are anagrams of each other.
Input: Two strings.
Output: Boolean (true/false).
Example:
Input: "anagram", "nagaram"
Output: true

## Valid Palindrome
Problem: Determine if a string is a palindrome, considering only alphanumeric characters and ignoring cases.
Input: A string.
Output: Boolean (true/false).
Example:
Input: "A man, a plan, a canal: Panama"
Output: true

## Valid Parentheses
Problem: Determine if the input string has valid matching parentheses.
Input: A string containing characters '(', ')', '{', '}', '[' and ']'.
Output: Boolean (true/false).
Example:
Input: "()[]{}"
Output: true

## Implement ATOI
Problem: Convert a string to a 32-bit signed integer (ASCII to integer).
Input: A string.
Output: 32-bit signed integer.
Example:
Input: "   -42"
Output: -42

## Implement STRSTR
Problem: Find the first occurrence of a needle in a haystack.
Input: Two strings: haystack and needle.
Output: Starting index of the first occurrence, or -1 if not found.
Example:
Input: haystack = "hello", needle = "ll"
Output: 2

## Z-Function
Problem: Computes the Z-array where Z[i] is the length of the longest common prefix between S and the suffix of S starting at i.
Input: A string S.
Output: An array of integers (Z-array).
Example:
Input: "aabcaabx"
Output: [0, 1, 0, 0, 3, 1, 0, 0]

## KMP Algorithm
Problem: String matching algorithm that uses a prefix function (LPS) to avoid redundant comparisons.
Input: text string and pattern string.
Output: Index/indices where pattern is found.
Example:
Input: text = "ababcabcabababd", pattern = "ababd"
Output: Found at index 10

## Rabin-Karp
Problem: String searching algorithm that uses hashing to find pattern strings in a text.
Input: text string and pattern string.
Output: Index/indices where pattern is found.
Example:
Input: text = "AABAACAADAABAABA", pattern = "AABA"
Output: Found at indices 0, 9, 12

## Minimum Window Substring
Problem: Find the minimum window in string S which will contain all the characters in string T.
Input: Two strings S and T.
Output: The minimum window substring.
Example:
Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"