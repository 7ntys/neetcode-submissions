func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	prefix[0] = 1
	for i:=1;i<len(nums);i++{
		prefix[i] = nums[i-1] * prefix[i-1]
	}

	suffix := make([]int, len(nums))
	suffix[len(suffix)-1] = 1
	for i:=len(nums)-2;i>=0;i--{
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	for i:=0;i<len(prefix);i++{
		prefix[i] *= suffix[i]
	}
	return prefix
}
