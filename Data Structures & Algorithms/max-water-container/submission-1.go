func maxArea(heights []int) int {
	l, r := 0, len(heights) - 1

	best := 0
	for l < r {
		length, width := min(heights[l], heights[r]), r - l

		best = max(best, computeArea(length, width))
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}

	}
	return best
}

func computeArea(length int, width int) int {
	return length * width
}
