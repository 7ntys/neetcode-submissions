
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

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// approach : 2 heaps : 
	// One min heap, one max heap (inversed sign)
	// always have a difference of & maximum => if more re equilibrate
	left, right := &MinHeap{}, &MinHeap{}
	heap.Init(left)
	heap.Init(right)

	for len(nums1) > 0 && len(nums2) > 0 {
		x := -1
		if nums1[0] >= nums2[0] {
			x = nums1[0]
			nums1 = nums1[1:]
		}  else {
			x = nums2[0]
			nums2 = nums2[1:]
		}

		processInt(x, left, right)
	}

	for i:=0;i<len(nums1);i++{
		processInt(nums1[i], left, right)
	}

	for i:=0;i<len(nums2);i++{
		processInt(nums2[i], left, right)
	}

	if (left.Len() + right.Len()) % 2 == 1 {
		return float64((*left)[0] * -1)
	}
	return (float64((*left)[0] * -1) + float64((*right)[0])) / 2.0
}

func processInt(x int, left *MinHeap, right *MinHeap) {
	// process x : 
	if left.Len() > 0 {
		// check if current number is higher than highest of left pane : 
		if x < (*left)[0] * -1 {
			heap.Push(left, x * -1)
		} else {
			heap.Push(right, x)
		}
	} else {
		heap.Push(left, -x)
	}

	// equilibrate : 
	for left.Len() > right.Len() + 1 {
		x := heap.Pop(left).(int) * -1
		heap.Push(right, x)
	}

	for right.Len() > left.Len() {
		x := heap.Pop(right).(int)
		heap.Push(left, x * -1)
	}
}
