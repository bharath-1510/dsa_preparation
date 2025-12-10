# Array Basics — Problem Statements & Examples

Below are short problem statements, input/output formats and a small example for each problem in this folder. These are concise references so you can quickly implement or test solutions.

## Find Largest Element in array
Problem: Given an array of n integers, return the largest element.
Input: n (size) then n integers.
Output: single integer — the largest value.
Example:
Input:
5
1 3 2 5 4
Output:
5

## Check Second Largest Element in Array
Problem: Given an array, find the second largest distinct element. If none exists, define behavior (e.g., return the largest or a sentinel).
Input: n then n integers.
Output: single integer — second largest distinct value.
Example:
Input:
5
2 3 6 6 1
Output:
3

## Rotate array by Left K
Problem: Rotate an array left by k positions (elements shifted to the left, wrap-around).
Input: n k then n integers.
Output: array after rotation.
Example:
Input:
5 2
1 2 3 4 5
Output:
3 4 5 1 2

## Rotate Array by Right K
Problem: Rotate an array right by k positions (elements shifted to the right, wrap-around).
Input: n k then n integers.
Output: array after rotation.
Example:
Input:
5 2
1 2 3 4 5
Output:
4 5 1 2 3

## Check Array sorted or not
Problem: Check whether an array is sorted in non-decreasing order.
Input: n then n integers.
Output: true or false (or 1/0).
Example:
Input:
4
1 2 2 3
Output:
true

## Remove array duplicates in sorted array
Problem: Given a sorted array, remove duplicates in-place and return the new length (or the deduped array).
Input: n then n sorted integers.
Output: new length and/or the deduped array.
Example:
Input:
5
1 1 2 2 3
Output:
3
1 2 3

## Move zero to end
Problem: Move all zeros in an array to the end while maintaining the relative order of non-zero elements.
Input: n then n integers.
Output: array after moving zeros.
Example:
Input:
5
0 1 0 3 12
Output:
1 3 12 0 0

## Array Linear Serach
Problem: Given an array and a target value, return the index of target or -1 if not found (linear search).
Input: n then n integers then target.
Output: index (0-based) or -1.
Example:
Input:
4
5 3 7 1
7
Output:
2

## Maximum Consecutive one's
Problem: Given a binary array, find the maximum number of consecutive 1s.
Input: n then n binary integers (0 or 1).
Output: single integer — longest run of 1s.
Example:
Input:
6
1 1 0 1 1 1
Output:
3

## Set Matrix Zeroes
Problem: Given an m x n matrix, if an element is 0 set its entire row and column to 0 (in-place if possible).
Input: m n then m rows of n integers.
Output: modified matrix.
Example:
Input:
2 2
1 0
3 4
Output:
0 0
3 0

## Pascal’s Triangle
Problem: Generate the first numRows of Pascal's triangle (each row is the binomial coefficients).
Input: single integer numRows.
Output: list of rows.
Example:
Input:
3
Output:
1
1 1
1 2 1

## Next Permutation
Problem: Rearrange numbers into the lexicographically next greater permutation. If not possible, rearrange into the lowest possible order (sorted ascending).
Input: n then n integers.
Output: modified array representing the next permutation.
Example:
Input:
3
1 3 2
Output:
2 1 3

## Kadane’s Algorithm
Problem: Given an integer array, find the contiguous subarray with the largest sum and return that sum.
Input: n then n integers.
Output: single integer — maximum subarray sum.
Example:
Input:
9
-2 1 -3 4 -1 2 1 -5 4
Output:
6   (subarray 4 -1 2 1)

## Sort Colors
Problem: Given an array with values 0,1,2 representing colors, sort them in-place so that elements of the same color are adjacent (Dutch National Flag problem).
Input: n then n integers (0/1/2).
Output: sorted array.
Example:
Input:
6
2 0 2 1 1 0
Output:
0 0 1 1 2 2

## Segregate 0s and 1s in an array
Problem: Given a binary array, move all 0s to one side and 1s to the other (stable not required unless specified).
Input: n then n binary integers.
Output: array with 0s and 1s segregated.
Example:
Input:
5
0 1 0 1 1
Output:
0 0 1 1 1

## Stock Buy & Sell
Problem: Given prices where prices[i] is price on day i, compute max profit by buying and selling once (buy before sell).
Input: n then n integers (prices).
Output: single integer — maximum profit (0 if no profit possible).
Example:
Input:
6
7 1 5 3 6 4
Output:
5   (buy at 1, sell at 6)

## Rotate Matrix
Problem: Rotate an NxN matrix by 90 degrees clockwise in-place.
Input: n then n rows of n integers.
Output: rotated matrix.
Example:
Input:
2
1 2
3 4
Output:
3 1
4 2

## Merge Overlapping Intervals
Problem: Given a list of intervals [start, end], merge all overlapping intervals and return the result.
Input: m then m pairs of integers.
Output: merged intervals.
Example:
Input:
3
1 3
2 6
8 10
Output:
[1,6] [8,10]

## Merge Sorted Arrays
Problem: Given two sorted arrays, merge them into a single sorted array.
Input: n m then n elements then m elements.
Output: merged sorted array.
Example:
Input:
3 3
1 3 5
2 4 6
Output:
1 2 3 4 5 6

## Find Missing Number
Problem: Given n distinct numbers in the range [0, n], find the missing number.
Input: n then n integers (one range value missing).
Output: the missing integer.
Example:
Input:
3
0 1 3
Output:
2

## Duplicate Number
Problem: Given an array containing n+1 integers where each integer is between 1 and n (inclusive), find the duplicate number.
Input: n+1 then values.
Output: the duplicate integer.
Example:
Input:
5
1 3 4 2 2
Output:
2

## Repeating + Missing Number
Problem: Given an array of size n containing numbers from 1..n with one number repeated and one missing, find both.
Input: n then n integers.
Output: repeated and missing numbers.
Example:
Input:
5
3 1 2 5 3
Output:
Repeated = 3, Missing = 4
