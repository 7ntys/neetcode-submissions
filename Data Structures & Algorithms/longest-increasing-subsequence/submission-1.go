func lengthOfLIS(nums []int) int {
    LIS := make(map[int]int)
	best := 0 

	for i:=0;i<len(nums);i++{
		best = max(best, dfs(nums, i, LIS))
	}
	return best
}

func dfs(nums []int, index int, LIS map[int]int) int {
	if val, ok := LIS[index]; ok {
		return val
	}

	best := 1

	for i:=index+1;i<len(nums);i++{
		if nums[index] < nums[i] {
			best = max(best, dfs(nums, i, LIS) + 1)
		}
	}
	LIS[index] = best
	return best
}