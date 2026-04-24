/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */


type MinHeap []Interval

func (h MinHeap) Len() int {return len(h)}
func (h MinHeap) Less(i,j int) bool {return h[i].end < h[j].end}
func (h MinHeap) Swap(i,j int) {h[i], h[j] = h[j], h[i]}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Interval))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minMeetingRooms(intervals []Interval) int {
	sort.Slice(intervals, func(i,j int) bool {
		return intervals[i].start < intervals[j].start
	})

	h := &MinHeap{}
	heap.Init(h)

	rooms := 0
	for _, inter := range intervals {
		for h.Len() > 0 && (*h)[0].end <= inter.start {
			heap.Pop(h)
		}
		heap.Push(h, inter)
		rooms = max(rooms, h.Len())
	}
	return rooms
}
