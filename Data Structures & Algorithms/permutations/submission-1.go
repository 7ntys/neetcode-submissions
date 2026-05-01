func permute(nums []int) [][]int {
	res := [][]int{}
	dfs(nums, &res, []int{}, make([]bool, len(nums)))
	return res
}

func dfs(nums []int, res *[][]int, path []int, chosen []bool){
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for i:=0;i<len(nums);i++{
		if chosen[i] {continue}
		path = append(path, nums[i])
		chosen[i] = true
		dfs(nums, res, path, chosen)
		chosen[i] = false
		path = path[:len(path)-1]
	}

}