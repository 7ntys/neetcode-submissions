func numIslands(grid [][]byte) int {
    islands := 0
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[i]);j++{
			if grid[i][j] == '1' {
				islands++
				dfs(grid, i, j)
			}
		}
	}
	return islands
}

func dfs(grid [][]byte, i int, j int) {
	for _, pos := range positions {
		x, y := i + pos[0], j + pos[1]

		if x < 0 || x >= len(grid) {continue}
		if y < 0 || y >= len(grid[x]) {continue}
		if grid[x][y] != '1' {continue}

		grid[x][y] = '0'
		dfs(grid, x, y)
	}
}

var positions = [][2]int{
	{1,0},
	{-1,0},
	{0,1},
	{0,-1},
}