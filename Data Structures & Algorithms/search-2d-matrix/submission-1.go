func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	l, r := 0, m * n - 1

	for l <= r {
		mid := (l + r) / 2

		row := mid / n
		col := mid % n

		if matrix[row][col] == target {return true}

		if matrix[row][col] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
