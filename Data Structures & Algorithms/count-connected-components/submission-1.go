func countComponents(n int, edges [][]int) int {
    parents := make([]int, n)
	for i:=0;i<n;i++{
		parents[i] = i
	}

	for _, e := range edges {
		a, b := e[0], e[1]
		Union(a, b, parents)
	}

	groups := make(map[int]struct{})
	for _, id := range parents {
		parent := Find(id, parents)
		if _, ok := groups[parent]; !ok {
			groups[parent] = struct{}{}
		}
	}

	return len(groups)
}

func Union(a int, b int, parents []int) {
	parentA, parentB := Find(a,parents), Find(b, parents)
	if parentA == parentB {return}
	parents[parentB] = parentA
}

func Find(a int, parents []int) int {
	if parents[a] != parents[parents[a]] {
		parents[a] = Find(parents[parents[a]], parents)
	}
	return parents[a]
}