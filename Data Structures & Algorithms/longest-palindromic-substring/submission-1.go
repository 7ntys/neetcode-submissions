func longestPalindrome(s string) string {
    bestL, bestR := 0, 0

	for i:=0;i<len(s);i++{

		l1, r1 := expand(s, i, i)
		l2, r2 := expand(s, i, i + 1)

		if bestR - bestL < r1 - l1 {
			bestL = l1
			bestR = r1
		}

		if bestR - bestL < r2 - l2 {
			bestL = l2
			bestR = r2
		}
	}
	return s[bestL:bestR+1]
}

func expand(s string, l int, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return l + 1, r - 1
}