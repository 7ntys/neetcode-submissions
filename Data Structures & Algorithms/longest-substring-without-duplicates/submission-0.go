func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	memo := make(map[byte]struct{})
	gap := 0
	for r < len(s) {
		if _, ok := memo[s[r]]; ok {
			for l < r && s[l] != s[r] {
				delete(memo, s[l])
				l++
			}
			// now s[l] = s[r]
			delete(memo, s[l])
			l++
		}
		memo[s[r]] = struct{}{}
		gap = max(gap, r - l + 1)
		r++
	}
	return gap
}
