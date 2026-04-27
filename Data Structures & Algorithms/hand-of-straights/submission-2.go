func isNStraightHand(hand []int, groupSize int) bool {
    data := make(map[int]int)
	if len(hand) % groupSize != 0 {return false}
	for _, x := range hand {
		data[x]++
	}

	sort.Ints(hand)

	for len(data) > 0 {
		fmt.Println(hand)
		x := hand[0]
		data[x]--
		if data[x] == 0 {
			delete(data, x)
		}
		// Can we make a group of size groupSize starting from here ? 
		for i:=1;i<groupSize;i++{
			if _, ok := data[x+1]; !ok {return false}
			x = x + 1
			data[x]--
			if data[x] == 0 {
				delete(data, x)
			}
		}

		for len(hand) > 0 && data[hand[0]] == 0 {
			hand = hand[1:]
		}
	}
	return true
}
