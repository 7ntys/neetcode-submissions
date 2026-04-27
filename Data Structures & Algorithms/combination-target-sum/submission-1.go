func combinationSum(nums []int, target int) [][]int {
    res := &[][]int{}
	sort.Ints(nums)
	backtrack(nums, target, 0, []int{}, res)
	return *res
}

func backtrack(nums []int, target int, index int, path []int, res *[][]int) {
	
	fmt.Println(target, path)
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i:=index;i<len(nums);i++{
		if target - nums[i] < 0 {
			break
		}

		// We can try this path 
		path = append(path, nums[i])
		backtrack(nums, target - nums[i], i, path, res)
		path = path[:len(path)-1]
	}
}