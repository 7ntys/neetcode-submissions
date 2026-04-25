func lengthOfLIS(nums []int) int {
    LIS := make(map[int]int)
	best := 0
	for i:=0;i<len(nums);i++{
		best = max(best, dfs(nums, i, LIS))
	}
	return best
}

func dfs(nums []int, i int, LIS map[int]int) int {
	if i == len(nums) {
		return 0
	}
	if val, ok := LIS[i]; ok {
		return val
	}
	best := 1
	for j:=i+1;j<len(nums);j++{
		if nums[j] > nums[i] {
			best = max(best, dfs(nums, j, LIS) + 1)
		}
	}
	LIS[i] = best
	return best
}