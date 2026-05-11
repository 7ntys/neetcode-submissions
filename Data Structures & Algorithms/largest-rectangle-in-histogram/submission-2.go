func largestRectangleArea(heights []int) int {
	s := []Node{}
	area := 0
	for i := range heights {
		offset := i
		for len(s) > 0 && s[len(s) - 1].height > heights[i] {
			n := s[len(s) - 1]
			s = s[:len(s) - 1]
			area = max(area, (i - n.index) * n.height)
			offset = min(offset, n.index)
		}

		s = append(s, Node{index: offset, height: heights[i]})
	}

	for len(s) > 0 {
		n := s[len(s) - 1]
		s = s[:len(s) - 1]
		area = max(area, (len(heights) - n.index) * n.height)
	}
	return area
}

type Node struct {
	index int
	height int
}