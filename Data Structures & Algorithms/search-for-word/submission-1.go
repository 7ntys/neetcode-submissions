func exist(board [][]byte, word string) bool {
	for i:=0;i<len(board);i++{
		for j:=0;j<len(board[i]);j++{
			if board[i][j] == word[0] {
				path := []byte{board[i][j]}
				board[i][j] = '_'
				if dfs(board, word, path, 1, i, j) {return true}
				board[i][j] = path[0]
			}
		}
	}
	return false
}

var positions = [][2]int{
    {1, 0},
    {-1, 0},
    {0, 1},
    {0, -1},
}

func dfs(board [][]byte, word string, path []byte, index int, i int, j int) bool {
	if len(path) == len(word) || index >= len(word) {return true}

	for _, pos := range positions {
		x, y := i + pos[0], j + pos[1]

		if x < 0 || x >= len(board) {continue}
		if y < 0 || y >= len(board[x]) {continue}

		if board[x][y] == word[index] {
			// See if its a correct path : 
			board[x][y] = '_'
			path = append(path, word[index])
			if dfs(board, word, path, index+1, x, y) {return true}
			// revert choice
			board[x][y] = word[index]
			path = path[:len(path)-1]
		}
	}
	return false
}