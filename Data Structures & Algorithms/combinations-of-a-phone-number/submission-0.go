func letterCombinations(digits string) []string {
	if len(digits) == 0 {return []string{}}
	store['2'] = []byte{'a','b','c'}
	store['3'] = []byte{'d','e','f'}
	store['4'] = []byte{'g','h','i'}
	store['5'] = []byte{'j','k','l'}
	store['6'] = []byte{'m','n','o'}
	store['7'] = []byte{'p','q','r','s'}
	store['8'] = []byte{'t','u','v'}
	store['9'] = []byte{'w','x','y','z'}
	res := &[]string{}
	dfs(digits, 0, []byte{}, res)
	return *res
}

var store = make(map[byte][]byte)


func dfs(digits string, i int, path []byte, res *[]string) {
	if i == len(digits) {
		temp := make([]byte, len(path))
		copy(temp,path)
		*res = append(*res, string(temp))
		return
	}

	digit := digits[i]

	for _, c := range store[digit] {
		path = append(path, c)
		dfs(digits, i + 1, path, res)
		path = path[:len(path)-1]
	}
} 