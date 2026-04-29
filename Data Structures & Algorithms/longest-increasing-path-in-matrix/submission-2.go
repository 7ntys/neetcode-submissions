func longestIncreasingPath(matrix [][]int) int {
    best := 1
	memo := make(map[[2]int]int)
	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[i]);j++{
			best = max(best, dfs(matrix, i, j, memo))
		}
	}
	return best
}

var positions = [][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func dfs(matrix [][]int, i int, j int, memo map[[2]int]int) int {
	if val, ok := memo[[2]int{i,j}]; ok {return val}
	best := 1
	for _, pos := range positions {
		x, y := i + pos[0], j + pos[1]

		if x < 0 || x > len(matrix) - 1 {continue}
		if y < 0 || y > len(matrix[x]) - 1 {continue}

		if matrix[i][j] >= matrix[x][y] {continue}
		best = max(best, dfs(matrix, x, y, memo) + 1)
	}
	memo[[2]int{i,j}] = best
	return best
}