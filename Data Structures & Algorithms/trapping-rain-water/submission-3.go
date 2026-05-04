func trap(height []int) int {
	n:= len(height)
	maxL := make([]int, n)
	maxL[0] = 0
	for i:=1;i<n;i++{
		maxL[i] = max(height[i-1], maxL[i-1])
	}

	maxR := make([]int, n)
	maxR[n-1] = 0
	for i:=n-2;i>=0;i--{
		maxR[i] = max(maxR[i+1], height[i+1])
	} 

	area := 0
	for i:=0;i<n;i++{
		diff := min(maxR[i], maxL[i]) - height[i]
		if diff > 0 {area += diff}
	}
	return area
}
