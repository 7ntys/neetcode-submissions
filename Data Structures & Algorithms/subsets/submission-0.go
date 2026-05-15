func subsets(nums []int) [][]int {
	res := [][]int{}
	dfs(nums, 0, []int{}, &res)
	return res
}

func dfs(nums []int, start int, path []int, res *[][]int) {
	temp := make([]int, len(path))
	copy(temp, path)
	*res = append(*res, temp)

	for i := start; i < len(nums); i++ {
		// chose this : 
		path = append(path, nums[i])
		dfs(nums, i + 1, path, res)
		path = path[:len(path) - 1]
	}
}