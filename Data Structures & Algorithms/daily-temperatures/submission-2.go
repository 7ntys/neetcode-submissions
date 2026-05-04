func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	res := make([]int, len(temperatures))

	for i := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			x := stack[len(stack) - 1]
			res[x] = i - x
			stack = stack[:len(stack) - 1]
		}

		stack = append(stack, i)
	}

	for _, x := range stack {
		res[x] = 0
	}
	return res
}
