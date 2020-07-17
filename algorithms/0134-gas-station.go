package algorithms

func canCompleteCircuit(gas []int, cost []int) int {
	// gen remain
	for i := 0; i < len(gas); i++ {
		gas[i] = gas[i] - cost[i]
	}

	index := 0
	for i := 1; i < len(gas); i++ {
		gas[i] += gas[i-1]
		if gas[i] < gas[index] {
			index = i
		}
	}
	if gas[len(gas)-1] < 0 {
		return -1
	}
	return (index + 1) % len(gas)
}
