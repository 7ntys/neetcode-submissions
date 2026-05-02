func rob(nums []int) int {
	memo := make(map[int]int)
    dfs(nums, 0, memo)
	return memo[0]
}

func dfs(nums []int, i int, memo map[int]int) int {
	if i >= len(nums) {
		return 0
	}
	if val, ok := memo[i]; ok {return val}
	
	// choose to take nums[i] : 
	memo[i] = max(memo[i], dfs(nums, i + 2, memo) + nums[i])

	// chose not to take : 
	memo[i] = max(memo[i], dfs(nums, i + 1, memo))

	return memo[i]
}