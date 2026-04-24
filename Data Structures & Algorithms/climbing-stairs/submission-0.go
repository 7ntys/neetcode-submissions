func climbStairs(n int) int {
    memo := make(map[int]int)
	memo[0] = 0
	memo[1] = 1
	memo[2] = 2
	return recurse(n, memo)
}

func recurse(n int, memo map[int]int) int {
	if val, ok := memo[n]; ok {return val}
	memo[n] = recurse(n-1, memo) + recurse(n-2, memo)
	return memo[n]
}