func canCompleteCircuit(gas []int, cost []int) int {
    fullTank := 0
	tank := 0
	index := 0
	for i:=0;i<len(gas);i++{
		diff := gas[i] - cost[i]
		fullTank+=diff
		tank+=diff
		if tank < 0 {
			tank = 0
			index = i + 1
		}
	}
	if fullTank < 0 {return -1}
	return index
}
