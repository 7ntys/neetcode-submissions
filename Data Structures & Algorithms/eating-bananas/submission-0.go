func minEatingSpeed(piles []int, h int) int {
	if h < len(piles) {return -1}
	n := 0
	for x := range piles {
		if piles[x] > n {n = piles[x]}
	}

	l, r := 1, n

	for l < r {
		mid := l + (r - l) / 2

		if feasible(piles, h, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func feasible(piles []int, h int, k int) bool {
	for x := range piles {
		h -= (piles[x] + k - 1) / k
		if h < 0 {return false}
	}
	return true
}
