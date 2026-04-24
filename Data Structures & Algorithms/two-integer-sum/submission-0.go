func twoSum(nums []int, target int) []int {
    memo := make(map[int]int)
	for i, num := range nums {
		if _, ok := memo[target - num]; ok {
			return []int{memo[target-num], i}
		}
		memo[num] = i
	}
	return []int{}
}
