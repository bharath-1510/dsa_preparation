# Matrix — Problem Statements & Examples

Below are concise problem statements, input/output formats and one example for each problem in this folder.

## Spiral Traversal
Problem: Given a matrix, return all elements in spiral order (clockwise from top-left).
Input: m rows, n columns, then m x n integers.
Output: list of elements in spiral order.
Example:
Input:
3 3
1 2 3
4 5 6
7 8 9
Output: 1 2 3 6 9 8 7 4 5

## Rotate Image
Problem: Rotate an n x n matrix by 90 degrees clockwise in-place.
Input: n, then n x n integers.
Output: rotated matrix.
Example:
Input:
3
1 2 3
4 5 6
7 8 9
Output:
7 4 1
8 5 2
9 6 3

## Word Search
Problem: Given a 2D board and a word, check if the word exists in the grid (can move horizontally or vertically, no cell reused).
Input: m rows, n columns, board, then word.
Output: true/false.
Example:
Input:
3 4
A B C E
S F C S
A D E E
"ABCCED"
Output: true

## Matrix Search
Problem: Search for a value in a matrix with each row and column sorted in ascending order.
Input: m rows, n columns, matrix, then target.
Output: true/false.
Example:
Input:
3 4
1 3 5 7
10 11 16 20
23 30 34 60
3
Output: true

## Max Rectangle in Matrix
Problem: Find the area of the largest rectangle containing only 1s in a binary matrix.
Input: m rows, n columns, matrix of 0/1.
Output: integer — area of largest rectangle.
Example:
Input:
4 4
0 1 1 0
1 1 1 1
1 1 1 1
1 1 0 0
Output: 8

## Row with Max 1s
Problem: Find the row with the maximum number of 1s in a binary matrix (rows sorted).
Input: m rows, n columns, matrix of 0/1.
Output: row index (0-based).
Example:
Input:
4 4
0 1 1 1
0 0 1 1
1 1 1 1
0 0 0 0
Output: 2

## Boolean Matrix
Problem: Given a binary matrix, if any cell is 1, set its entire row and column to 1.
Input: m rows, n columns, matrix of 0/1.
Output: modified matrix.
Example:
Input:
2 2
1 0
0 0
Output:
1 1
1 0

## Islands in 2D grid
Problem: Count the number of islands (groups of connected 1s) in a 2D binary grid (connected horizontally/vertically).
Input: m rows, n columns, matrix of 0/1.
Output: integer — number of islands.
Example:
Input:
4 5
1 1 0 0 0
1 1 0 0 0
0 0 1 0 0
0 0 0 1 1
Output: 3

## Shortest path in matrix
Problem: Find the shortest path from top-left to bottom-right in a binary matrix (0 = open, 1 = blocked), moving in 8 directions.
Input: n rows, n columns, matrix of 0/1.
Output: integer — length of shortest path or -1 if not possible.
Example:
Input:
3 3
0 1 0
1 0 1
0 0 0
Output: 4

## Flood Fill
Problem: Given a matrix, a start cell, and a new color, fill all connected cells of the same color as the start cell with the new color (4 directions).
Input: m rows, n columns, matrix, start row, start col, new color.
Output: modified matrix.
Example:
Input:
3 3
1 1 1
1 1 0
1 0 1
1 1 2
Output:
2 2 2
2 2 0
2 0 1