func carFleet(target int, position []int, speed []int) int {
	cars := []Car{}

	for i:=0;i<len(position);i++{
		cars = append(cars, Car{position: position[i], speed: speed[i]})
	}

	sort.Slice(cars, func(i,j int) bool {
		return cars[i].position > cars[j].position
	})

	groups := 0
	max := 0.0
	for _, c := range cars {
		// time taken by car 'c' to achieve destination
		time := float64(target - c.position) / float64(c.speed)
		if time > max {
			max = time
			groups++
		}
	}
	return groups
}

type Car struct {
	position int
	speed int
}