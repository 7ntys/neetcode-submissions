func findItinerary(tickets [][]string) []string {
    adj := make(map[string][]string)

	for _, t := range tickets {
		from, to := t[0], t[1]

		if _, ok := adj[from]; !ok {
			adj[from] = []string{}
		}

		adj[from] = append(adj[from], to)
	}

	for src := range adj {
		sort.Sort(sort.Reverse(sort.StringSlice(adj[src])))
	}

	res := []string{}
	dfs(adj, "JFK", &res)

	i, j := 0, len(res) - 1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}

func dfs(adj map[string][]string, curr string, res *[]string) {
	for len(adj[curr]) > 0 {
		n := len(adj[curr])
		next := adj[curr][n-1]
		adj[curr] = adj[curr][:n-1]

		dfs(adj, next, res)
	}
	*res = append(*res, curr)
}



