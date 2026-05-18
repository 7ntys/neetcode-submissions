func jump(nums []int) int {
    jumps := 0
	furthest := 0
	currentEnd := 0

	i := 0
	for i < len(nums) - 1 {
		x := nums[i]
		furthest = max(furthest, i + x)

		if currentEnd == i {
			currentEnd = furthest
			jumps++
		}
		i++
	}
	return jumps
}
