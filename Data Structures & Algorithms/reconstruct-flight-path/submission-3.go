func findItinerary(tickets [][]string) []string {
    adj := make(map[string][]*Airport)

	for _, t := range tickets {
		from, to := t[0], t[1]

		if _, ok := adj[from]; !ok {
			adj[from] = []*Airport{}
		}

		airport := &Airport{id: to, visited: false}
		adj[from] = append(adj[from], airport)
	}

	for from := range adj {
		sort.Slice(adj[from], func(i,j int) bool {
			return adj[from][i].id < adj[from][j].id
		})
	}
	
	res := []string{}
	dfs(adj, "JFK", []string{"JFK"}, len(tickets) + 1, &res)
	return res
}

func dfs(adj map[string][]*Airport, curr string, path []string, numberOfFlight int, res *[]string) bool {
	if len(path) == numberOfFlight {
		temp := make([]string, len(path))
		copy(temp, path)
		*res = temp
		return true
	}

	for _, a := range adj[curr] {
		if a.visited {continue}

		// take this path : 
		path = append(path, a.id)
		a.visited = true
		if dfs(adj, a.id, path, numberOfFlight, res) {return true}
		a.visited = false
		path = path[:len(path)-1]
	}
	return false
}


type Airport struct {
	id string
	visited bool
}