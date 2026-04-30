func countSubstrings(s string) int {
    total := 0

	for i:=0;i<len(s);i++{
		total += expand(s, i, i)
		total += expand(s, i, i + 1)
	}
	return total
}

func expand(s string, l int, r int) int {
	counter := 0

	for l > -1 && r < len(s) && s[l] == s[r] {
		l--
		r++
		counter++
	}
	return counter
}