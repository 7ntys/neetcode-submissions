func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Print(nums)
	res := [][]int{}
	for i:=0;i<len(nums)-2;i++{
		if nums[i] > 0 {break}
		if i > 0 && nums[i] == nums[i-1] {continue}
		l, r := i+1, len(nums) - 1

		for l < r {
			if nums[i] + nums[l] + nums[r] > 0 {
				r--
			} else if nums[i] + nums[l] + nums[r] < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				for l < len(nums) && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return res
}
