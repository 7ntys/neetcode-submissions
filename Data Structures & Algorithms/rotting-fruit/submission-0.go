func orangesRotting(grid [][]int) int {
    q := []Point{}
	rotten := 0
	oranges := 0
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[i]);j++{
			if grid[i][j] != 0 {oranges++}
			if grid[i][j] == 2 {
				rotten++
				q = append(q, Point{i:i, j:j})
			}
		}
	}
	time := 0
	positions := [][2]int{
		{1,0},
		{-1,0},
		{0,1},
		{0,-1},
	}
	for len(q) > 0 && rotten < oranges {
		size := len(q)
		for i:=0;i<size;i++{
			p := q[i]

			for _, pos := range positions {
				x, y := p.i + pos[0], p.j + pos[1]
				if x < 0 || x >= len(grid) {continue}
				if y < 0 || y >= len(grid[x]) {continue}
				if grid[x][y] != 1 {continue}

				grid[x][y] = 2
				rotten++
				q = append(q, Point{i:x, j:y})
			}
		}
		q = q[size:]
		time++
	}
	if rotten == oranges {return time}
	return -1	
}

type Point struct {
	i int
	j int
}