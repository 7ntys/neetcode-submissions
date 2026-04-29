func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {return []int{}}

	res := []int{}
	top, left := 0, 0
	bottom, right := len(matrix) - 1, len(matrix[0]) - 1

	for top <= bottom && left <= right {
		// Left to right
		for col := left; col <= right; col++{
			res = append(res, matrix[top][col])
		}
		top++

		// Up to down
		for row := top; row <= bottom; row++{
			res = append(res, matrix[row][right])
		}
		right--

		// Right to left
		if bottom >= top {
			for col := right; col >= left; col-- {
				res = append(res, matrix[bottom][col])
			}
		}
		bottom--

		// Down to up
		if left <= right {
			for row := bottom; row >= top; row-- {
				res = append(res, matrix[row][left])
			}
		}
		left++
	}
	return res
}
