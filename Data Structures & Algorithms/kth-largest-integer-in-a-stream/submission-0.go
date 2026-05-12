
type MinHeap []int

func (h MinHeap) Len() int {return len(h)}
func (h MinHeap) Less(i,j int) bool {return h[i] < h[j]}
func (h MinHeap) Swap(i,j int) {h[i], h[j] = h[j], h[i]}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type KthLargest struct {
    h *MinHeap
	k int
}


func Constructor(k int, nums []int) KthLargest {
    h := &MinHeap{}
	heap.Init(h)

	for _, x := range nums {
		heap.Push(h, x)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return KthLargest{h: h, k: k}

}


func (this *KthLargest) Add(val int) int {
    heap.Push(this.h, val)
	if this.h.Len() > this.k {
		heap.Pop(this.h)
	}
	return (*this.h)[0]
}
