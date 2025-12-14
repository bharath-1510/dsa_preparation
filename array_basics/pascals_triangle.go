package arraybasics

// pascalsTriangle generates the first n rows of Pascal's triangle.
// Time Complexity: O(n^2) — building each row requires O(row_length) work.
// Space Complexity: O(n^2) — result stores O(n^2) numbers.
// Best achievable: Time = O(n^2), Space = O(n^2) — storing all rows requires O(n^2); streaming single rows can reduce output space.
func pascalsTriangle(n int) [][]int {
	var res [][]int
	if n == 0 {
		return [][]int{}
	}

	res = append(res, []int{1})
	for i := 1; i < n; i++ {
		var temp []int
		for j := 0; j <= i; j++ {
			a := getValueFromAbove(i-1, j-1, res)
			b := getValueFromAbove(i-1, j, res)
			temp = append(temp, (a + b))
		}
		res = append(res, temp)
	}
	return res
}

func getValueFromAbove(line, pos int, res [][]int) int {
	n := len(res[line])
	if pos >= n || pos < 0 {
		return 0
	}
	return res[line][pos]
}
