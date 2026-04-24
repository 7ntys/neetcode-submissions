func pacificAtlantic(heights [][]int) [][]int {
    pac := make([][]int, len(heights))
	atl := make([][]int, len(heights))
	for i:=0;i<len(heights);i++{
		pac[i] = make([]int, len(heights[i]))
		atl[i] = make([]int, len(heights[i]))
	}
	pacQ := []Point{}
	atlQ := []Point{}
	for i:=0;i<len(heights);i++{
		pac[i][0] = 1 // meaning we can reach
		atl[i][len(heights[i]) - 1] = 1
		pacQ = append(pacQ, Point{i: i, j: 0})
		atlQ = append(atlQ, Point{i: i, j: len(heights[i]) - 1})
	}
	for j:=0;j<len(heights[0]);j++{
		pac[0][j] = 1 // meaning we can reach
		atl[len(heights) - 1][j] = 1
		pacQ = append(pacQ, Point{i: 0, j: j})
		atlQ = append(atlQ, Point{i: len(heights) - 1, j:j})
	}
	bfs(heights, &pac, pacQ)
	bfs(heights, &atl, atlQ)

	res := [][]int{}
	for i:=0;i<len(heights);i++{
		for j:=0;j<len(heights[i]);j++ {
			if pac[i][j] == 1 && atl[i][j] == 1 {
				res = append(res, []int{i,j})
			}
		}
	}
	return res
}

type Point struct {
	i int
	j int
}

func bfs(heights [][]int, ocean *[][]int, q []Point) {
	positions := [][2]int{
		{1,0},
		{-1,0},
		{0,1},
		{0,-1},
	}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		for _, pos := range positions {
			i, j := p.i + pos[0], p.j + pos[1]

			if i < 0 || i >= len(heights) {continue}
			if j < 0 || j >= len(heights[i]) {continue}
			if (*ocean)[i][j] != 0 {continue}
			if heights[p.i][p.j] > heights[i][j] {continue}
			(*ocean)[i][j] = 1
			q = append(q, Point{i:i, j:j})
		}
	}
}