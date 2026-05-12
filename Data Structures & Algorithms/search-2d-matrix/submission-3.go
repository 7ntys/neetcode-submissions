func searchMatrix(matrix [][]int, target int) bool {
	r := searchRow(matrix, target)
	if r < 0 || r >= len(matrix) {return false}
	c := searchCol(matrix, target, r)
	if c == -1 {return false}
	return matrix[r][c] == target
}

func searchCol(matrix [][]int, target int, right int) int {
	n := len(matrix[right])

	l, r := 0, n - 1

	for l <= r {
		mid := l + (r - l) / 2

		if matrix[right][mid] == target {return mid}

		if matrix[right][mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func searchRow(matrix [][]int, target int) int {
	m := len(matrix) - 1

	l, r := 0, m

	for l <= r {
		mid := l + (r - l) / 2

		if matrix[mid][0] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}