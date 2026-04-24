func countBits(n int) []int {
	output := []int{}
	for i:=0;i<=n;i++{
		binary := convertToBinary(i)
		output = append(output, countOnes(binary))
	}
	return output
}

func convertToBinary(n int) []int {
	// from 1024 -> 0
	// 1024 = 0100 0000 0000
	start := 1024
	res := []int{}
	for i:=0;i<11;i++{
		if n - start >= 0 {
			n -= start
			res = append(res, 1)
		} else {
			res = append(res, 0)
		}
		if start == 1 {start = 0}
		start /= 2
	}
	fmt.Printf("Converted %d to : [%d]\n",n,res)
	return res
}

func countOnes(array []int) int {
	ones := 0
	for _, val := range array {
		if val == 1 {ones++}
	}
	return ones
}