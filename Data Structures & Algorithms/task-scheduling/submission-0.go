type MaxHeap []Task

type Task struct {
	symbol byte
	count int
	timeout int
}

func (h MaxHeap) Len() int {return len(h)}
func (h MaxHeap) Less(i,j int) bool {return h[i].count > h[j].count}
func (h MaxHeap) Swap(i,j int) {h[i], h[j] = h[j], h[i]}

func (h *MaxHeap) Push(x interface{}){
	*h = append(*h, x.(Task))
}

func (h * MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
	counter := make(map[byte]int)

	for _, t := range tasks {
		counter[t]++
	}
	h := &MaxHeap{}
	heap.Init(h)
	for k, v := range counter {
		t := Task{symbol: k, count: v}
		heap.Push(h,t)
	}

	q := []Task{}
	cycle := 0
	for h.Len() > 0 || len(q) > 0 {		
		cycle++
		if h.Len() > 0 {
			t := heap.Pop(h).(Task)
			t.count--
			if t.count > 0 {
				t.timeout = cycle + n
				q = append(q, t)
			}
		}
		
		for len(q) > 0 && q[0].timeout == cycle {
			heap.Push(h, q[0])
			q = q[1:]
		}
	}
	return cycle
}
