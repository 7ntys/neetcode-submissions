func longestPalindrome(s string) string {
    if len(s) == 0 {return ""}
	size := 1
	curr := string(s[0])

	for i:=0;i<len(s);i++{
		
		pair, unpair := getPairPalindromes(s, i), getOddPalindromes(s, i)

		if size < len(pair) {
			size = len(pair)
			curr = pair
		}
		if size < len(unpair) {
			size = len(unpair)
			curr = unpair
		}
	}
	return curr
}

func getPairPalindromes(s string, i int) string {
	l, r := i, i + 1 
	
	var left strings.Builder
	var right strings.Builder
	
	for l > - 1 && r < len(s) && s[l] == s[r] {
		left.WriteByte(s[l])
		right.WriteByte(s[r])
		l--
		r++
	}
	return reverseBytes(left.String()) + right.String()
} 
func getOddPalindromes(s string, i int) string {
	l, r := i-1, i+1

	var left strings.Builder
	var right strings.Builder
	left.WriteByte(s[i])
	for l > -1 && r < len(s) && s[l] == s[r] {
		left.WriteByte(s[l])
		right.WriteByte(s[r])
		l--
		r++ 
	}
	return reverseBytes(left.String()) + right.String()
}	

func reverseBytes(s string) string {
	b := []byte(s)
	i, j := 0, len(b) - 1

	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}





