const INF = int(1e9)
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
	for i:=0;i<len(dp);i++{
		dp[i] = INF
	}
	dp[amount] = 0
	for i:=amount;i>=0;i--{
		if dp[i] != INF {
			for _, c := range coins {
				if i - c >= 0 && dp[i-c] > dp[i] + 1 {
					dp[i-c] = dp[i] + 1
				}
			}
		}
	}
	if dp[0] == INF {return -1}
	return dp[0]
}
