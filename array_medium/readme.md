# Array Medium — Problem Statements & Examples

Below are concise problem statements, input/output formats and one example for each problem in this folder.

## Inversion Count
Problem: Count the number of inversions in an array — pairs (i, j) such that i < j and arr[i] > arr[j].
Input: n then n integers.
Output: single integer — inversion count.
Example:
Input: 5
2 4 1 3 5
Output: 3  (inversions: (2,1), (4,1), (4,3))

## Find the number that appears once, and other numbers twice
Problem: In an array where every element appears twice except one element which appears once, find that single element.
Input: n then n integers (n is odd).
Output: the single integer that appears once.
Example:
Input: 5
2 3 2 4 3
Output: 4

## Search in Rotated Sorted Array
Problem: Given a rotated sorted array (ascending) and a target, return the index of target or -1 if not found.
Input: n then n integers (rotated sorted) then target.
Output: index (0-based) or -1.
Example:
Input: 7
4 5 6 7 0 1 2
0
Output: 4

## 2-Sum
Problem: Given an array and a target sum, find a pair of indices (i, j) such that arr[i] + arr[j] = target. Return any one valid pair.
Input: n then n integers then target.
Output: pair of indices (0-based) or indication none.
Example:
Input: 4
2 7 11 15
9
Output: 0 1

## 4-Sum
Problem: Given an array and a target, find all unique quadruplets [a,b,c,d] that sum to target.
Input: n then n integers then target.
Output: list of quadruplets (values) without duplicates.
Example:
Input: 6
1 0 -1 0 -2 2
0
Output: [-2,-1,1,2] [-2,0,0,2] [-1,0,0,1]

## Longest Consecutive Sequence
Problem: Given an unsorted array, find the length of the longest sequence of consecutive integers.
Input: n then n integers.
Output: single integer — length of longest consecutive sequence.
Example:
Input: 8
100 4 200 1 3 2 2 5
Output: 5  (sequence 1,2,3,4,5)

## Largest Subarray with 0 Sum
Problem: Given an array of integers, find the length of the largest contiguous subarray whose sum is 0.
Input: n then n integers.
Output: single integer — maximum length.
Example:
Input: 6
15 -2 2 -8 1 7
Output: 5  (subarray from index 1 to 5 sums to 0)

## Longest subarray with given sum K (positives)
Problem: For an array of positive integers, find the longest contiguous subarray with sum equal to K.
Input: n then n positive integers then K.
Output: length of longest subarray or 0 if none.
Example:
Input: 6
1 2 3 7 5 2
12
Output: 3  (subarray 2,3,7)

## Longest subarray with sum K (Positives + Negatives)
Problem: For an array containing positives and negatives, find the longest contiguous subarray whose sum equals K.
Input: n then n integers then K.
Output: length of longest subarray or 0 if none.
Example:
Input: 5
10 5 2 7 1
15
Output: 2  (subarray 10,5)

## Count Subarrays with Given XOR
Problem: Count subarrays where XOR of elements equals a given value X.
Input: n then n integers then X.
Output: single integer — count of subarrays.
Example:
Input: 4
4 2 2 6
6
Output: 2  (subarrays [4,2] and [6])

## Count special triplets
Problem: Count triplets (i, j, k) satisfying XOR(i..j-1) == XOR(j..k) for all 0 ≤ i < j ≤ k < n. This is the usual "special triplets" formulation.
Input: n then n integers.
Output: single integer — number of special triplets.
Example:
Input: 4
1 2 3 1
Output: 1  (triplet (0,1,3))

## Longest Substring without Repeat
Problem: Given a string, find the length of the longest substring without repeating characters.
Input: single string s.
Output: integer — length of longest substring with unique chars.
Example:
Input: "abcabcbb"
Output: 3  (substring "abc")

## Number of Subarrays with Given Sum (k)
Problem: Count the number of contiguous subarrays whose sum equals k. Array may contain negatives.
Input: n then n integers then k.
Output: integer — count of such subarrays.
Example:
Input: 5
1 1 1 2 1
3
Output: 3  (subarrays: [1,1,1], [1,2], [2,1])

## Majority Element I
Problem: Given an array, find the element that appears more than ⌊n/2⌋ times (guaranteed to exist).
Input: n then n integers.
Output: the majority element.
Example:
Input: 5
3 3 4 2 3
Output: 3

## Majority Element II
Problem: Given an array, find all elements that appear more than ⌊n/3⌋ times.
Input: n then n integers.
Output: list of elements (0–2 elements).
Example:
Input: 7
3 2 3 2 2 3 3
Output: 3 2

## Subarray with Given Product
Problem: Given an array of positive integers and a target P, find a contiguous subarray whose product equals P (or return indices of any one such subarray).
Input: n then n positive integers then P.
Output: indices of subarray (start,end) 0-based or -1 if none.
Example:
Input: 5
2 3 4 2 6
12
Output: 1 2  (subarray 3*4 = 12)

## Trapping Rain Water
Problem: Given heights of bars, compute how much water can be trapped after raining.
Input: n then n non-negative integers (height array).
Output: single integer — total trapped water units.
Example:
Input: 6
0 1 0 2 1 0
Output: 2
