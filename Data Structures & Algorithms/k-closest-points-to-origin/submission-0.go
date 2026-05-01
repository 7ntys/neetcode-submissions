
type MinHeap []Point

type Point struct {
	x int
	y int
	distance float64
}

func (h MinHeap) Len() int {return len(h)}
func (h MinHeap) Less(i,j int) bool {
	return h[i].distance > h[j].distance
}
func (h MinHeap) Swap(i,j int) {h[i], h[j] = h[j], h[i]}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (p Point) computeDistance() float64 {
	return math.Sqrt(float64(p.x * p.x) + float64(p.y * p.y))
}

func kClosest(points [][]int, k int) [][]int {
	h := &MinHeap{}
	heap.Init(h)

	for _, point := range points {
		p := Point{x: point[0], y: point[1]}
		p.distance = p.computeDistance()
		fmt.Println(p)
		heap.Push(h, p)

		for h.Len() > k {
			heap.Pop(h)
		}
	}

	res := [][]int{}

	for h.Len() > 0 {
		p := heap.Pop(h).(Point)
		res = append(res, []int{p.x,p.y})
	}
	return res
}
