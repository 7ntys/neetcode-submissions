func rob(nums []int) int {
	if len(nums) == 0 {return 0}
	if len(nums) == 1 {return nums[0]}
    return max(robHouses(nums[1:]), robHouses(nums[:len(nums)-1]))
}

func robHouses(nums []int) int {
	rob1, rob2 := 0, 0
	for _, n := range nums {
		temp := max(rob2, rob1 + n)
		rob1 = rob2
		rob2 = temp
	}
	return rob2
}