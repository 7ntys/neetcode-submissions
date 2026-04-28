func insert(intervals [][]int, newInterval []int) [][]int {
    res := [][]int{}

	i := 0
	// deal with non overlapping from left side
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	// deal with overlapping : 
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
		i++
	}
	res = append(res, newInterval)

	// deal with non overlapping from right side
	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}

	return res
}
