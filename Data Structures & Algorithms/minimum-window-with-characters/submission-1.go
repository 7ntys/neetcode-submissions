func minWindow(s string, t string) string {
    res := ""

	tracker := make(map[byte]int)

	for x := range t {
		tracker[t[x]]++
	}

	l, r := 0, 0
	validChar := 0
	for r < len(s) {
		char := s[r]
		
		if _, ok := tracker[char]; ok {
			tracker[char]--
			if tracker[char] == 0 {
				// we have the right number of char : 
				validChar++
			}
		}

		for validChar == len(tracker) {
			// All of the char of t is present in substring : shrink 
			if r - l + 1 < len(res) || res == "" {
				res = s[l:r+1]
			}

			// make sure to increment tracker
			if _, ok := tracker[s[l]]; ok {
				tracker[s[l]]++
				if tracker[s[l]] == 1 {
					// One less valid char in window : 
					validChar--
				}
			}
			l++
		}
	
		r++
	}
	return res
}
