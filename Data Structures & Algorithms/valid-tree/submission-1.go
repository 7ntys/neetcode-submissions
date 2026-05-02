func validTree(n int, edges [][]int) bool {
    parents := make([]int, n)
    for i := range parents {
        parents[i] = i
    }

    for _, e := range edges {
        a, b := e[0], e[1]

        cycle := Union(a, b, parents)
        if !cycle {return false}
    }
    parents[0] = Find(0, parents)
    for i:=1;i<len(parents);i++ {
        parents[i] = Find(i, parents)
        if parents[i] != parents[i-1] {return false}
    }
    return true
}

func Union(a int, b int, parents []int) bool {
    parentA, parentB := Find(a, parents), Find(b, parents)

    if parentA == parentB {
        return false
    }
    parents[parentB] = parentA
    return true
}

func Find(a int, parents []int) int {
    if a != parents[a] {
        parents[a] = Find(parents[a], parents)
    }
    return parents[a]
}   


