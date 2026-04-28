func partition(s string) [][]string {
	res := &[][]string{}
	dfs(s, 0, []string{}, res)
	return *res
}

func dfs(s string, start int, path []string, res *[][]string) {
	if start == len(s) {
		temp := make([]string, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for i:=start;i<len(s);i++{
		if isPalindrome(s[start:i+1]) {
			path = append(path, s[start:i+1])
			dfs(s, i + 1, path, res)
			path = path[:len(path) - 1]
		}
	}
	return 
}

func isPalindrome(s string) bool {
	if len(s) == 0 {return false}
	l, r := 0, len(s) - 1 

	for l <= r && s[l] == s[r] {
		l++
		r--
	}
	return l >= r
}