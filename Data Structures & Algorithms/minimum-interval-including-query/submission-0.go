type interval struct {
	left int
	right int
	length int
	index int
}

func minInterval(intervals [][]int, queries []int) []int {
    array := []interval{}
	for i, inter := range intervals {
		array = append(array, interval{left: inter[0], right: inter[1], length: inter[1] - inter[0] + 1, index: i})
	}

	sort.Slice(array, func(i,j int) bool {
		return array[i].length < array[j].length
	})

	db := make(map[int]int)
	for _, inter := range array {
		for i:=inter.left;i<=inter.right;i++{
			if _, ok := db[i]; !ok {
				db[i] = inter.length
			}
		}
	}

	res := []int{}

	for _, q := range queries {
		if _, ok := db[q]; !ok {
			res = append(res, -1)
		} else {
			res = append(res, db[q])
		}
	}
	return res
}
