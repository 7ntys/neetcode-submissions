func isPalindrome(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	replace := strings.NewReplacer(
		" ", "",
		",", "",
		".", "",
		":", "",
		";", "",
		"!", "",
		"?", "",
		"'", "",
		"\"", "",
		"-", "",
		"_", "",
		"(", "",
		")", "",
	)

	s = replace.Replace(s)
	l, r := 0, len(s) - 1
	for l <= r && s[l] == s[r] {
		l++
		r--
	}

	return l > r
}
