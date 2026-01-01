# Binary Search on Answer — Problem Statements & Examples

Below are concise problem statements, input/output formats and one example for each problem in this folder.

## Allocate Books
Problem: Given an array of book pages and students, allocate books such that the maximum pages assigned to a student is minimized. Each student gets a contiguous sequence of books.
Input: n books, m students, array of pages.
Output: Minimum possible maximum pages.
Example:
Input: 4 2
12 34 67 90
Output: 113

## Painter’s Partition
Problem: Given boards of different lengths and painters, partition boards among painters to minimize the time to paint all boards. Each painter paints contiguous boards.
Input: n boards, k painters, array of board lengths.
Output: Minimum time to paint all boards.
Example:
Input: 5 2
10 20 30 40 50
Output: 90

## Aggressive Cows
Problem: Place cows in stalls so that the minimum distance between any two cows is maximized.
Input: n stalls, k cows, array of stall positions.
Output: Largest minimum distance.
Example:
Input: 5 3
1 2 8 4 9
Output: 3

## Median of Two Sorted Arrays
Problem: Find the median of two sorted arrays.
Input: Two sorted arrays.
Output: Median value (float).
Example:
Input: [1, 3], [2]
Output: 2.0

## Kth Element of Two Sorted Arrays
Problem: Find the k-th smallest element in the union of two sorted arrays.
Input: Two sorted arrays, integer k.
Output: k-th smallest element.
Example:
Input: [2, 3, 6, 7, 9], [1, 4, 8, 10], k=5
Output: 6

## Single Element in Sorted Array
Problem: In a sorted array where every element appears twice except one, find the single element.
Input: Sorted array.
Output: The single element.
Example:
Input: [1,1,2,3,3,4,4,8,8]
Output: 2

## Floor & Ceil
Problem: Given a sorted array and a value x, find the floor (greatest <= x) and ceil (smallest >= x).
Input: Sorted array, value x.
Output: Floor and ceil values.
Example:
Input: [1,2,8,10,10,12,19], x=5
Output: Floor=2, Ceil=8

## Koko Eating Bananas
Problem: Koko can eat up to k bananas per hour. Given piles of bananas, find the minimum k so she can eat all bananas in h hours.
Input: Array of piles, integer h.
Output: Minimum integer k.
Example:
Input: [3,6,7,11], h=8
Output: 4

## Time to Eat Bananas
Problem: Given piles of bananas and hours h, find the minimum eating speed to finish all bananas in h hours.
Input: Array of piles, integer h.
Output: Minimum integer speed.
Example:
Input: [30,11,23,4,20], h=5
Output: 30

## Split Array Largest Sum
Problem: Split an array into m non-empty subarrays to minimize the largest sum among these subarrays.
Input: Array of integers, integer m.
Output: Minimum largest sum.
Example:
Input: [7,2,5,10,8], m=2
Output: 18

## Kth Missing Positive Number
Problem: Given a sorted array and integer k, find the k-th missing positive number.
Input: Sorted array, integer k.
Output: k-th missing positive number.
Example:
Input: [2,3,4,7,11], k=5
Output: 9

## Find Peak Element
Problem: Find a peak element in an array (greater than neighbors).
Input: Array of integers.
Output: Index of a peak element.
Example:
Input: [1,2,1,3,5,6,4]
Output: 1 or 5