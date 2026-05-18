func jump(nums []int) int {
    jumps := 0
	furthest := 0
	currentEnd := 0
	
	for i, x := range nums {
		if i == len(nums) - 1 {break}
		furthest = max(furthest, i + x)

		if currentEnd == i {
			currentEnd = furthest
			jumps++
		}
	}
	return jumps
}
