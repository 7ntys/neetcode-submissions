var positions = [][2]int{
	{1,0},
	{-1,0},
	{0,1},
	{0,-1},
}

func islandsAndTreasure(grid [][]int) {
    q := []Point{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				q = append(q, Point{i: i, j: j})
			}
		}
	}

	dist := 0
	for len(q) > 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			p := q[i]

			for _, pos := range positions {
				x, y := p.i + pos[0], p.j + pos[1]

				if x < 0 || x >= len(grid) {continue}
				if y < 0 || y >= len(grid[x]) {continue}
				if grid[x][y] != 2147483647 {continue}

				grid[x][y] = dist + 1
				q = append(q, Point{i: x, j: y})
			}
		}
		dist++
		q = q[size:]
	}
}


type Point struct {
	i int
	j int
}