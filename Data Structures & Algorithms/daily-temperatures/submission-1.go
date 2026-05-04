func dailyTemperatures(temperatures []int) []int {
	l, r := 0, 1
	res := make([]int, len(temperatures))
	for l < len(temperatures) - 1{
		if temperatures[r] > temperatures[l] {
			res[l] = r - l
			l++
			r = l + 1
		} else if r == len(temperatures) - 1 {
			res[l] = 0
			l++
			r = l + 1
		} else {
			r++
		}
	}
	res[len(res) - 1] = 0
	return res
}
