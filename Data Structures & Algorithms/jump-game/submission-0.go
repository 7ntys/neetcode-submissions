func canJump(nums []int) bool {
	q := []int{0}
	visited := make([]bool, len(nums))
	visited[0] = true
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		
		for j:=i+1;j<=i+nums[i];j++{
			if j >= len(nums) - 1 {return true}
			q = append(q, j)
			visited[j] = true
		}
	}
	return visited[len(nums) - 1]
}
