func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) + len(s2) != len(s3) {return false}
	return dfs(s1, s2, s3, 0, 0, 0, make(map[[2]int]bool))
}

func dfs(s1 string, s2 string, s3 string, i int, j int, k int, memo map[[2]int]bool) bool {
	if k == len(s3) {
		return true
	}
	key := [2]int{i,j}
	if val, ok := memo[key]; ok {return val}

	// try by having a s1 subgroup : 
	if i < len(s1) && s1[i] == s3[k] {
		memo[key] = dfs(s1, s2, s3, i + 1, j, k + 1, memo)
		if memo[key] {return true}
	}
	// continue with s2 subgroup : 
	if j < len(s2) && s2[j] == s3[k] {
		memo[key] = dfs(s1, s2, s3, i, j + 1, k + 1, memo) 
		if memo[key] {return true}
	}

	memo[[2]int{i,j}] = false
	return false
}