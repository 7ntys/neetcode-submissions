func canFinish(n int, prerequisites [][]int) bool {
    in := make([]int, n)
	
	deps := make([][]int, n)
	for i:=0;i<n;i++{
		deps[i] = []int{}
	}

	for _, preq := range prerequisites {
		a, b := preq[0], preq[1]

		deps[b] = append(deps[b], a)
		in[a]++
	}

	q := []int{}
	for i, count := range in {
		if count == 0 {q = append(q, i)}
	}

	processed := 0
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		processed++

		for _, dep := range deps[x] {
			in[dep]--
			if in[dep] == 0 {q = append(q, dep)}
		}
	}
	return processed == n
}
