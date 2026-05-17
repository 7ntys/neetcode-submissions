func longestCommonSubsequence(text1 string, text2 string) int {
    return dfs(text1, text2, 0, 0, make(map[[2]int]int))
}

func dfs(text1 string, text2 string, i int, j int, memo map[[2]int]int) int{
	if i == len(text1) || j == len(text2) {return 0}

	key := [2]int{i,j}

	if v, ok := memo[key]; ok {return v}

	if text1[i] == text2[j] {
		memo[key] = 1 + dfs(text1, text2, i + 1, j + 1, memo)
	} else {
		memo[key] = max(dfs(text1, text2, i + 1, j, memo), dfs(text1, text2, i, j + 1, memo))
	}
	return memo[key]
}