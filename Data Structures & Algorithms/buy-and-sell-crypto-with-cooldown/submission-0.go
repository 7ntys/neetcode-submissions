func maxProfit(prices []int) int {
    return dfs(prices, 0, 0, 0, make(map[[3]int]int))
}

func dfs(prices []int, i int, cooldown int, hold int, memo map[[3]int]int) int {
	if i == len(prices) {return 0}

	key := [3]int{i, cooldown, hold}

	if v, ok := memo[key]; ok {return v}

	// Are we on cooldown ? 
	if cooldown > 0 {
		// Only choice is to wait : 
		memo[key] = dfs(prices, i + 1, 0, hold, memo)
	} else {
		// Are we holding ?
		if hold > 0 {
			// We can sell : 
			memo[key] = max(memo[key], prices[i] + dfs(prices, i + 1, 1, 0, memo))
		} else {
			// We can buy :
			memo[key] = max(memo[key], -1 * prices[i] + dfs(prices, i + 1, 0, 1, memo))
		}
		// We can just wait :
		memo[key] = max(memo[key], dfs(prices, i + 1, 0, hold, memo))
	}
	return memo[key]
}