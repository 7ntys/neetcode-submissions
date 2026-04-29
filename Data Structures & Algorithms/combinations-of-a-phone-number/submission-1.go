var store = map[byte][]byte{
    '2': {'a', 'b', 'c'},
    '3': {'d', 'e', 'f'},
    '4': {'g', 'h', 'i'},
    '5': {'j', 'k', 'l'},
    '6': {'m', 'n', 'o'},
    '7': {'p', 'q', 'r', 's'},
    '8': {'t', 'u', 'v'},
    '9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }

    res := []string{}
    dfs(digits, 0, []byte{}, &res)
    return res
}

func dfs(digits string, i int, path []byte, res *[]string) {
    if i == len(digits) {
        *res = append(*res, string(path))
        return
    }

    digit := digits[i]

    for _, c := range store[digit] {
        path = append(path, c)
        dfs(digits, i+1, path, res)
        path = path[:len(path)-1]
    }
}