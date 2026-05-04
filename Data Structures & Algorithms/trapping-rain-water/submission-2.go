func trap(height []int) int {
	maxL := make([]int, len(height))
	maxR := make([]int, len(height))
	
	maxL[0] = 0
	for i:=1;i<len(height);i++{
		maxL[i] = max(maxL[i-1], height[i-1])
	}

	maxR[len(maxR)-1] = 0
	for i:=len(height)-2;i>=0;i--{
		maxR[i] = max(maxR[i+1], height[i+1])
	}

	area := 0
	for i:=0;i<len(height);i++{
		diff := min(maxL[i], maxR[i]) - height[i]
		if diff > 0 {
			area += diff
		}
	}
	return area
}
