func minCostClimbingStairs(cost []int) int {
    dp := make([]int, len(cost) + 1)
	dp[len(cost)] = 0
	dp[len(cost) - 1] = cost[len(cost) - 1]

	for i:=len(cost)-2;i>=0;i--{
		dp[i] = min(dp[i+1], dp[i+2]) + cost[i]
	}
	return min(dp[0], dp[1])
}
