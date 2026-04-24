func ladderLength(beginWord string, endWord string, wordList []string) int {
    db := make(map[string][]string)
	for _, word := range wordList{
		// Go through all chars and replace one by one by '*'

		for i:=0;i<len(word);i++{
			temp := word[:i] + "*" + word[i+1:]
			db[temp] = append(db[temp], word)
		}
	}

	q := []string{beginWord}
	steps := 0
	for len(q) > 0 {
		size := len(q)
		steps++
		for i:=0;i<size;i++{
			w := q[i]
			if w == endWord {return steps}
			for j:=0;j<len(w);j++{
				x := w[:j] + "*" + w[j+1:]

				if _, ok := db[x]; !ok {continue}
				q = append(q, db[x]...)
				delete(db, x)
			}
		}
		q = q[size:]
	}
	return 0
}
