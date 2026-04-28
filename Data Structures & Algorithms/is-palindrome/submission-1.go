func isPalindrome(s string) bool {
	l, r := 0, len(s) - 1

	for l <= r {
		a, b := rune(s[l]), rune(s[r])

		for l < r && (!unicode.IsLetter(a) && !unicode.IsDigit(a)) {
			l++
			a = rune(s[l])
		}

		for l < r && (!unicode.IsLetter(b) && !unicode.IsDigit(b)) {
			r--
			b = rune(s[r])
		}

		if unicode.ToLower(a) != unicode.ToLower(b) {break}
		l++
		r--
	}
	return l > r
}
