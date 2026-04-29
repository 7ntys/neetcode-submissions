func minDistance(word1 string, word2 string) int {
    return dfs(word1, word2, 0, 0, make(map[[2]int]int))
}

func dfs(word1 string, word2 string, i int, j int, memo map[[2]int]int) int {
	if i == len(word1) {return len(word2) - j}
	if j == len(word2) {return len(word1) - i}
	
	key := [2]int{i,j}

	if val, ok := memo[key] ;ok {return val}

	if word1[i] == word2[j] {
		memo[key] = dfs(word1, word2, i + 1, j + 1, memo)
		return memo[key]
	}

	// Insert : 
	memo[key] = 1 + dfs(word1, word2, i, j+1, memo)

	// Remove :
	memo[key] = min(memo[key], 1 + dfs(word1, word2, i + 1, j, memo))

	// Replace : 
	memo[key] = min(memo[key], 1 + dfs(word1, word2, i + 1, j + 1, memo))

	return memo[key]
}