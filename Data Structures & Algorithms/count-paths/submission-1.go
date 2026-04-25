func uniquePaths(m int, n int) int {
    dp := make([][]int, m + 1)
	for i:=0;i<=m;i++{
		dp[i] = make([]int, n + 1)
	}

	i := m - 1
	j := n - 1
	dp[i][j] = 1
	j--
	for i >= 0 {
		for j >= 0 {
			dp[i][j] = dp[i+1][j] + dp[i][j+1]
			j--
		}
		i--
		j = n - 1
	}
	return dp[0][0]
}
