func change(amount int, coins []int) int {
    return dfs(amount, coins, 0, make(map[[2]int]int))
}

func dfs(a int, coins []int, i int, memo map[[2]int]int) int {
	if a == 0 {return 1}
	if a < 0 {return 0}
	if i == len(coins) {return 0}

	key := [2]int{a, i}
	if v, ok := memo[key]; ok {return v}

	// take coin coins[i] : 
	memo[key] += dfs(a - coins[i], coins, i, memo)

	// Don't take coin coins[i]
	memo[key] += dfs(a, coins, i + 1, memo)

	return memo[key]
}